package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/semantic"
	"github.com/gabivlj/candice/internals/undomap"
	"github.com/gabivlj/candice/pkg/logger"
	"github.com/gabivlj/candice/pkg/random"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"strings"
)

var zero value.Value = constant.NewInt(types.I32, 0)
var one value.Value = constant.NewInt(types.I32, 1)

type Compiler struct {
	errors []error
	m      *ir.Module
	blocks []*ir.Block

	// Current defined types
	types map[string]*Type

	// Defined values
	globalBuiltinDefinitions map[string]value.Value

	concatenationUtilsLoaded bool
	builtins                 map[string]func(*Compiler, *ast.BuiltinCall) value.Value
	globalVariables          map[string]*Value

	variables *undomap.UndoMap[string, value.Value]

	currentFunction *ir.Func

	// Flag to indicate caller that this value shouldn't be loaded into memory
	// This flag will be set to false again once loadIfPointer is called.
	doNotLoadIntoMemory bool
	doNotAllocate       bool

	currentBreakLeaveBlock     *ir.Block
	currentContinueEscapeBlock *ir.Block
	context                    *semantic.Semantic
	modules                    map[string]*Compiler
	compiledModules            map[string]*Compiler

	eventHandler func(Event)
}

func New(context *semantic.Semantic, parent ...*Compiler) *Compiler {
	var m *ir.Module
	var builtins map[string]func(*Compiler, *ast.BuiltinCall) value.Value
	var globalBuiltinDefinitions map[string]value.Value
	var globalVariables map[string]*Value
	var compiledModules map[string]*Compiler

	if len(parent) > 0 {
		// we need previous module to add llvm IR here.
		m = parent[0].m

		// define it for extern funcs so we don't redefine them
		globalVariables = parent[0].globalVariables

		// builtin functions so we don't redefine later
		builtins = parent[0].builtins

		globalBuiltinDefinitions = parent[0].globalBuiltinDefinitions

		// When importing we might import an already compiled module, but with another name,
		// let's remember those!
		compiledModules = parent[0].compiledModules

	} else {
		m = ir.NewModule()
		globalVariables = map[string]*Value{}
		builtins = map[string]func(*Compiler, *ast.BuiltinCall) value.Value{}
		globalBuiltinDefinitions = map[string]value.Value{}
		compiledModules = map[string]*Compiler{}
	}

	c := &Compiler{
		m:                        m,
		blocks:                   []*ir.Block{},
		globalBuiltinDefinitions: globalBuiltinDefinitions,
		builtins:                 builtins,
		types:                    map[string]*Type{},
		variables:                undomap.New[string, value.Value](),
		globalVariables:          globalVariables,
		currentFunction:          nil,
		context:                  context,
		modules:                  map[string]*Compiler{},
		compiledModules:          compiledModules,
		eventHandler:             func(e Event) {},
	}

	c.variables.Add("<>", nil)

	if len(parent) > 0 {
		// necessary because we need references to types defined before importing.
		c.modules["_parent"] = parent[0]
		c.types = parent[0].types
	}

	c.initializeBuiltinLib()
	// define structs and functions
	c.compileStructTypes(context.Root.Statements)
	c.compileFunctionTypes(context.Root.Statements)
	return c
}

// A small context stores variables stored by the current scope defined by if and for statements.
func (c *Compiler) createSmallContext() {
	c.variables.Add("<>", nil)
}

func (c *Compiler) removeCurrentSmallContext() {
	for s, _ := c.variables.Pop(); s != "<>"; s, _ = c.variables.Pop() {
	}
}

func (c *Compiler) pushBlock(block *ir.Block) {
	c.blocks = append(c.blocks, block)
	c.createSmallContext()
}

func (c *Compiler) block() *ir.Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *Compiler) popBlock() *ir.Block {
	if len(c.blocks) == 0 {
		return nil
	}
	b := c.block()
	c.removeCurrentSmallContext()
	c.blocks = c.blocks[:len(c.blocks)-1]
	return b
}

func (c *Compiler) initializeBuiltinLib() {
	if _, ok := c.builtins["print"]; ok {
		return
	}

	c.builtins["add_compiler_flag"] = func(c *Compiler, addCompilerFlag *ast.BuiltinCall) value.Value {
		for _, param := range addCompilerFlag.Parameters {
			str := c.compileConstantExpression(param).(*constant.CharArray)
			c.eventHandler(Event{
				Kind: AddFlags,
				Data: string(str.X[:len(str.X)-1]),
			})
		}

		return constant.NewUndef(types.Void)
	}

	c.builtins["free"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		ptr := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		returnedValue := c.block().NewCall(c.free(), c.block().NewBitCast(ptr, types.I8Ptr))
		return returnedValue
	}

	c.builtins["asm"] = func(c *Compiler, bc *ast.BuiltinCall) value.Value {
		expressions := make([]value.Value, 0, len(bc.Parameters)-1)
		for _, parameter := range bc.Parameters[1:] {
			expressions = append(expressions, c.loadIfPointer(c.compileExpression(parameter)))
		}
		return c.asm(c.ToLLVMType(bc.TypeParameters[0]), bc.Parameters[0].(*ast.StringLiteral).Value, expressions...)
	}

	c.builtins["realloc"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		typeParameter := call.GetType()
		toReturnType := c.ToLLVMType(typeParameter)
		ptr := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		length := c.loadIfPointer(c.compileExpression(call.Parameters[1]))
		length = c.handleIntegerCast(types.I64, length)
		totalSize := c.block().NewMul(length, constant.NewInt(types.I64, typeParameter.SizeOf()))
		returnedValue := c.block().NewCall(c.realloc(), c.block().NewBitCast(ptr, types.I8Ptr), totalSize)
		castedValue := c.block().NewBitCast(returnedValue, toReturnType)
		alloca := c.block().NewAlloca(castedValue.Type())
		c.block().NewStore(castedValue, alloca)
		return alloca
	}

	c.builtins["sizeof"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		return constant.NewInt(types.I32, call.TypeParameters[0].SizeOf())
	}

	c.builtins["print"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		expressions := make([]value.Value, len(call.Parameters)+1)
		for i := 0; i < len(call.Parameters); i++ {
			expressions[i+1] = c.loadIfPointer(c.compileExpression(call.Parameters[i]))
		}
		constantString := strings.Builder{}
		for i := range call.Parameters {
			t := expressions[i+1].Type()
			if i != 0 {
				constantString.WriteString(" " + c.getFormatString(expressions, t, call, i))
				continue
			}

			constantString.WriteString(c.getFormatString(expressions, t, call, i))
		}
		s := constantString.String()
		expressions[0] = c.createString(s)
		c.block().NewCall(c.printf(), expressions...)
		return constant.NewUndef(types.Void)
	}

	c.builtins["unreachable"] = func(c *Compiler, _ *ast.BuiltinCall) value.Value {
		c.block().NewUnreachable()
		return constant.NewUndef(types.Void)
	}

	// alloc accepts one type parameter, and how many you want to allocate
	c.builtins["alloc"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		typeParameter := call.TypeParameters[0]
		toReturnType := types.NewPointer(c.ToLLVMType(typeParameter))
		length := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		length = c.handleIntegerCast(types.I64, length)
		totalSize := c.block().NewMul(length, constant.NewInt(types.I64, typeParameter.SizeOf()))
		returnedValue := c.block().NewCall(c.malloc(), totalSize)
		castedValue := c.block().NewBitCast(returnedValue, toReturnType)
		alloca := c.block().NewAlloca(castedValue.Type())
		c.block().NewStore(castedValue, alloca)
		return alloca
	}

	c.builtins["cast"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		return c.handleCast(call)
	}

}

/// Public methods for the compiler

func (c *Compiler) GenerateExecutable() error {
	err := GenerateExecutable(c.m, "exec")
	return err
}

// GenerateExecutableCXX creates an executable from llvm bitcode, CXX needs to be able to compile llvm
func (c *Compiler) GenerateExecutableCXX(output string, cxx string, flags []string) error {
	defer func() {
		os.Remove(".intermediate_output.ll")
	}()
	fd, _ := os.Create(".intermediate_output.ll")
	_, _ = c.m.WriteTo(fd)
	endFlags := []string{".intermediate_output.ll", "-o", output}
	endFlags = append(endFlags, flags...)
	cmd := exec.Command(cxx, endFlags...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error compiling with %s:\n"+stdout.String()+"\n status: "+err.Error(), cxx)
	}

	return nil
}

func (c *Compiler) GenerateExecutableExperimental(output string, cxx string, flags []string, optimized bool, link bool) error {
	objectOutput := output
	pathOutput, err := GenerateObjectLLVM(c.m, objectOutput, optimized)
	if err != nil {
		return err
	}

	if !link {
		return err
	}

	command := append(flags, pathOutput)
	command = append(command, "-o", output)
	// command = append(command, "-O3")
	cmd := exec.Command(cxx, command...)
	outputBuffer := bytes.Buffer{}
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &outputBuffer
	err = cmd.Run()
	if err != nil {
		return errors.New("error compiling " + strings.Join(command, " ") + " :\n" + outputBuffer.String())
	}
	return err
}

// Execute generates and executes the executable
func (c *Compiler) Execute() ([]byte, error) {
	err := GenerateExecutable(c.m, "exec")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("./exec")
	b := bytes.Buffer{}
	cmd.Stdout = &b
	e := cmd.Run()
	return bytes.Trim(b.Bytes(), " "), e
}

func (c *Compiler) CompileWithEventHandler(tree ast.Node, eventHandler func(Event)) {
	c.eventHandler = eventHandler
	c.Compile(tree)
}

// Compile compiles the entire ast
// It makes weak type checks, it will assume that the returned types
// are right. Usually you would want to semantically check the tree before
// calling this.
func (c *Compiler) Compile(tree ast.Node) {
	defer func() {
		// Reset state
		c.doNotLoadIntoMemory = false
		c.doNotAllocate = false
	}()

	if c.currentFunction != nil && len(c.currentFunction.Blocks) > 0 && c.currentFunction.Blocks[len(c.currentFunction.Blocks)-1].Term != nil {
		return
	}

	switch t := tree.(type) {
	case *ast.MultipleDeclarationStatement:
		{
			c.compileMultipleDeclarationStatement(t)
		}

	case *ast.SwitchStatement:
		{
			c.compileSwitchStatement(t)
		}

	case *ast.MacroBlock:
		{
			for _, statement := range t.Block.Statements {
				c.Compile(statement)
			}

			return
		}

	case *ast.Block:
		{
			block := c.currentFunction.NewBlock("block." + random.RandomString(10))
			c.block().NewBr(block)
			currentBlock := c.compileBlock(t, block)
			if currentBlock.Term == nil {
				c.blocks[len(c.blocks)-1] = currentBlock
			}
			return
		}

	case *ast.ImportStatement:
		{
			moduleName := t.Name
			module := c.context.GetModule(moduleName)
			if existingCompiler, ok := c.compiledModules[module.Root.ID]; ok {
				c.compiledModules[module.Root.ID] = existingCompiler
				c.modules[moduleName] = existingCompiler
				return
			}

			localCompiler := New(module, c)
			localCompiler.CompileWithEventHandler(module.Root, c.eventHandler)
			c.compiledModules[module.Root.ID] = localCompiler
			c.modules[moduleName] = localCompiler
			return
		}

	case *ast.BreakStatement:
		{
			c.block().NewBr(c.currentBreakLeaveBlock)
			return
		}

	case *ast.ContinueStatement:
		{
			c.block().NewBr(c.currentContinueEscapeBlock)
			return
		}

	case *ast.ExpressionStatement:
		{
			c.compileExpression(t.Expression)
			return
		}

	case *ast.GenericTypeDefinition:
		{
			return
		}

	case *ast.StructStatement:
		{
			// c.compileStruct(t)
			return
		}

	case *ast.DeclarationStatement:
		{
			c.compileDeclaration(t)
			return
		}

	case *ast.AssignmentStatement:
		{
			c.compileAssignment(t)
			return
		}

	case *ast.Program:
		{
			for _, statement := range t.Statements {
				c.Compile(statement)
			}

			if len(c.blocks) > 0 && c.block().Term == nil {
				c.block().NewRet(constant.NewInt(types.I32, 0))
			}

			//_ = c.popBlock()
			return
		}

	case *ast.FunctionDeclarationStatement:
		{
			c.compileFunctionDeclaration(t.FunctionType.Name, t)
			return
		}

	case *ast.ReturnStatement:
		{
			c.compileReturn(t)
			return
		}

	case *ast.IfStatement:
		{
			c.compileIf(t)
			return
		}

	case *ast.ForStatement:
		{
			c.compileFor(t)
			return
		}

	case *ast.ExternStatement:
		{
			c.compileExternFunc(t)
			return
		}

	case *ast.TypeDefinition:
		{

			return
		}
	}
}

func (c *Compiler) compileTypeDefinition(t *ast.TypeDefinition) {
	llvmType := c.ToLLVMType(t.Type)
	c.types[t.Name] = &Type{candiceType: t.Type, llvmType: llvmType}
}

func (c *Compiler) compileExternFunc(externFunc *ast.ExternStatement) {
	funcType := externFunc.Type.(*ctypes.Function)

	// Check if the variable has been defined anywhere in LLVM
	if f, ok := c.globalVariables[funcType.ExternalName]; ok {
		c.globalVariables[funcType.Name] = f
		return
	}

	if f, ok := c.globalBuiltinDefinitions[funcType.ExternalName]; ok {
		funk := &Value{Value: f, Type: funcType}
		c.globalVariables[funcType.Name] = funk
		return
	}

	returnType := c.ToLLVMType(funcType.Return)
	params := []*ir.Param{}
	for _, parameter := range funcType.Parameters {
		parameterType := c.ToLLVMType(parameter)
		params = append(params, ir.NewParam("", parameterType))
	}

	f := c.m.NewFunc(funcType.ExternalName, returnType, params...)
	f.CallingConv = enum.CallingConvC
	if funcType.InfiniteParameters {
		f.Sig.Variadic = true
	}

	funk := &Value{Value: f, Type: funcType}
	// Define the variable to local module facing references
	c.globalVariables[funcType.Name] = funk
	// Define the variable to C ABI facing code so the rest of modules
	// can know if this variable has been defined
	c.globalVariables[funcType.ExternalName] = funk
	// Define the variable for builtin functions in case that they are trying to
	// define this variable. For example @alloc wants to define malloc but it has been
	// defined on an extern funcs
	c.globalBuiltinDefinitions[funcType.ExternalName] = funk.Value
}

func (c *Compiler) compileReturn(ret *ast.ReturnStatement) {
	if ret.Expression == nil {
		if c.currentFunction.Name() == "main" {
			c.block().NewRet(zero)
			return
		}
		c.block().NewRet(nil)
		return
	}
	toReturn := c.compileExpression(ret.Expression)
	toReturnLoaded := c.loadIfPointer(toReturn)
	c.block().NewRet(toReturnLoaded)
}

// defines function type without compiling its body
func (c *Compiler) compileFunctionType(name string, funk ast.Function) *ir.Func {
	if fn, ok := c.globalVariables[name]; ok {
		return fn.Value.(*ir.Func)
	}

	functionType := funk.GetFunctionType()
	// Declare params LLVM IR
	params := make([]*ir.Param, 0, len(functionType.Parameters))
	for i, param := range functionType.Parameters {
		t := c.ToLLVMType(param)
		name := functionType.Names[i]
		params = append(params, ir.NewParam(name, t))
	}

	if functionType.IsMainFunction() && functionType.Return == ctypes.VoidType || functionType.Return == nil {
		functionType.Return = ctypes.I32
	}

	toReturnType := c.ToLLVMType(functionType.Return)
	// Declare llvmFunction
	llvmFunction := c.m.NewFunc(functionType.Name, toReturnType, params...)

	if functionType.RedefineWithOriginalName {
		llvmFunctionExtern := c.m.NewFunc(functionType.ExternalName, c.ToLLVMType(functionType.Return), params...)
		llvmFunctionExtern.CallingConv = enum.CallingConvC
		c.globalVariables[functionType.ExternalName] = &Value{
			Value: llvmFunctionExtern,
			Type:  functionType,
		}
	}

	// Create function
	c.globalVariables[name] = &Value{
		Value: llvmFunction,
		Type:  functionType,
	}

	return llvmFunction
}

// compiles entire function
func (c *Compiler) compileFunctionDeclaration(name string, funk ast.Function, variablesToAdd ...NamedValue) {
	llvmFunction := c.compileFunctionType(name, funk)

	// Create a main block to the function
	c.pushBlock(llvmFunction.NewBlock(funk.GetFunctionType().Name))

	// Declare parameters IR
	for _, param := range llvmFunction.Params {
		allocatedParameter := c.block().NewAlloca(param.Type())
		c.block().NewStore(param, allocatedParameter)
		c.declare(param.Name(), allocatedParameter)
	}

	for _, variable := range variablesToAdd {
		c.declare(variable.Name, variable.Value)
	}

	// Set it as current function
	prevFunction := c.currentFunction
	c.currentFunction = llvmFunction

	// Compile block
	for _, statement := range funk.GetBlock().Statements {
		c.Compile(statement)
	}

	lastBlock := c.currentFunction.Blocks[len(c.currentFunction.Blocks)-1]
	// If return hasn't been declared, declare a void return
	if lastBlock.Term == nil {
		if funk.GetFunctionType().IsMainFunction() {
			lastBlock.NewRet(zero)
		} else if funk.GetFunctionType().Return == ctypes.VoidType {
			lastBlock.NewRet(nil)
		} else {
			lastBlock.NewUnreachable()
		}
	}

	if funk.GetFunctionType().RedefineWithOriginalName {
		c.compileFunctionRedeclaration(c.currentFunction.Blocks, funk)
	}

	// Pop block, stack and restore current function
	c.popBlock()
	c.currentFunction = prevFunction
}

func (c *Compiler) compileFunctionRedeclaration(otherFunctionBlocks []*ir.Block, funk ast.Function) {
	llvmFunction := c.retrieveVariable(funk.GetFunctionType().ExternalName).(*ir.Func)
	llvmFunction.Blocks = otherFunctionBlocks
}

func (c *Compiler) compileIf(ifStatement *ast.IfStatement) {
	strandedBlocks := []*ir.Block{}
	block := c.currentFunction.NewBlock("if.then." + random.RandomString(10))
	strandedBlocks = append(strandedBlocks, c.compileBlock(ifStatement.Block, block))
	blockElseToJump := c.currentFunction.NewBlock("if.else." + random.RandomString(10))
	blocks := []*ir.Block{block}
	conditions := []ast.Expression{ifStatement.Condition}
	if ifStatement.Else != nil {
		blockElse := c.compileBlock(ifStatement.Else, blockElseToJump)
		strandedBlocks = append(strandedBlocks, blockElse)
	}

	for _, elseIf := range ifStatement.ElseIfs {
		currentBlock := c.currentFunction.NewBlock("elseif.then." + random.RandomString(10))
		// Append to stranded blocks that are either unreachable or need to keep going with their execution
		strandedBlocks = append(strandedBlocks, c.compileBlock(elseIf.Block, currentBlock))
		blocks = append(blocks, currentBlock)
		conditions = append(conditions, elseIf.Condition)
	}

	lastJumpToCondition := c.block()
	for i, currentBlock := range blocks {

		jumpToNextCondition := blockElseToJump
		if i+1 < len(blocks) {
			jumpToNextCondition = c.currentFunction.NewBlock("leave." + random.RandomString(10))
		}

		c.pushBlock(lastJumpToCondition)

		// If last block that needs to make a condition
		// is true jump to this block, else jump to the next one, which would be the next jump condition or
		// the else
		condition := c.compileExpression(conditions[i])
		conditionLoaded := c.loadIfPointer(condition)
		// Recharge block because compiling expressions can change current blocks
		lastJumpToCondition = c.popBlock()
		lastJumpToCondition.NewCondBr(c.toBool(conditionLoaded), currentBlock, jumpToNextCondition)
		lastJumpToCondition = jumpToNextCondition
	}

	leaveBlock := c.currentFunction.NewBlock("lastLeave." + random.RandomString(10))

	if lastJumpToCondition.Term == nil {
		lastJumpToCondition.NewBr(leaveBlock)
	}

	// override initial block
	c.blocks[len(c.blocks)-1] = leaveBlock

	for _, currentBlock := range blocks {
		if currentBlock.Term == nil {
			currentBlock.NewBr(leaveBlock)
		}
	}

	for _, currentBlock := range strandedBlocks {
		if currentBlock.Term == nil {
			currentBlock.NewBr(leaveBlock)
		}
	}

}

func (c *Compiler) compileFor(forLoop *ast.ForStatement) {
	// Exit point
	leave := ir.NewBlock("leave." + random.RandomString(10))
	blockDeclaration := c.currentFunction.NewBlock("for.declaration." + random.RandomString(10))

	c.block().NewBr(blockDeclaration)
	c.pushBlock(blockDeclaration)
	if forLoop.InitializerStatement != nil {
		c.Compile(forLoop.InitializerStatement)
	}

	condition := c.currentFunction.NewBlock("for.condition." + random.RandomString(10))
	conditionValueFirst := c.toBool(c.loadIfPointer(c.compileExpression(forLoop.Condition)))
	mainLoop := c.currentFunction.NewBlock("for.block." + random.RandomString(10))
	update := c.currentFunction.NewBlock("for.update." + random.RandomString(10))

	// jumps to main loop
	c.block().NewCondBr(conditionValueFirst, mainLoop, leave)

	previousBreak, previousContinue := c.currentBreakLeaveBlock, c.currentContinueEscapeBlock
	c.currentBreakLeaveBlock, c.currentContinueEscapeBlock = leave, update
	// compile main loop
	possibleNewBlock := c.compileBlock(forLoop.Block, mainLoop)
	c.currentBreakLeaveBlock, c.currentContinueEscapeBlock = previousBreak, previousContinue

	// compile update statement
	c.compileBlock(&ast.Block{Statements: []ast.Statement{forLoop.Operation}}, update)

	// jump to the update statement
	if possibleNewBlock.Term == nil {
		possibleNewBlock.NewBr(update)
	}

	// go to the condition again
	update.NewBr(condition)

	// Compile condition block
	c.pushBlock(condition)
	valueCondition := c.toBool(c.loadIfPointer(c.compileExpression(forLoop.Condition)))
	condition = c.popBlock()
	condition.NewCondBr(valueCondition, mainLoop, leave)

	// pop for loop block
	c.popBlock()
	c.currentFunction.Blocks = append(c.currentFunction.Blocks, leave)
	leave.Parent = c.currentFunction
	// normalize current block to the leave block
	c.blocks[len(c.blocks)-1] = leave
}

func (c *Compiler) compileBlock(block *ast.Block, blockIR *ir.Block) *ir.Block {
	c.pushBlock(blockIR)
	for _, statement := range block.Statements {
		c.Compile(statement)
	}
	return c.popBlock()
}

func (c *Compiler) compileDeclaration(decl *ast.DeclarationStatement) {
	t := c.ToLLVMType(decl.Type)
	var valueCompiled value.Value
	if decl.Constant {
		// If the declaration was a constant we don't need to allocate later on
		c.doNotAllocate = true
		valueCompiled = c.compileConstantExpression(decl.Expression)
		// If it is a global just compile it as a constant
	} else if c.currentFunction == nil {
		valueCompiled = c.compileConstantExpression(decl.Expression)
	} else {
		valueCompiled = c.compileExpression(decl.Expression)
	}

	if c.currentFunction == nil {
		c.doNotAllocate = false
		if _, isConstant := valueCompiled.(constant.Constant); isConstant {
			c.globalVariables[decl.Name] = &Value{Value: valueCompiled, Type: decl.Type, Constant: decl.Constant}
		} else {
			c.exitErrorExpression("Candice does not support non-constant global expressions and operations yet.", decl.Expression)
		}
		return
	}

	var val value.Value
	if !c.doNotAllocate {
		val = c.block().NewAlloca(t)
		valueCompiled = c.loadIfPointer(valueCompiled)
		c.block().NewStore(valueCompiled, c.bitcastIfUnion(decl.Type, val, types.NewPointer(valueCompiled.Type())))
	} else {
		c.doNotAllocate = false
		val = valueCompiled
	}

	c.declare(decl.Name, val)
}

func (c *Compiler) bitcastIfUnion(t ctypes.Type, toBeBitcasted value.Value, unionType types.Type) value.Value {
	possibleUnion := c.retrieveInnerAnonymousAndUnwrap(t)

	if _, isUnion := possibleUnion.(*ctypes.Union); isUnion {
		return c.block().NewBitCast(toBeBitcasted, unionType)
	}

	return toBeBitcasted
}

func (c *Compiler) declare(name string, value value.Value) {
	c.variables.Add(name, value)
}

func (c *Compiler) compileStruct(strukt *ast.StructStatement) {
	c.compileType(strukt.Type.Name, strukt.Type)
}

func (c *Compiler) compileUnion(union *ast.UnionStatement) {
	c.compileType(union.Type.Name, union.Type)
}

func (c *Compiler) compileType(name string, ct ctypes.Type) {
	t := c.ToLLVMType(ct)
	c.types[name] = &Type{
		llvmType:    c.m.NewTypeDef(name, t),
		candiceType: ct,
	}
}

func (c *Compiler) compileExpression(expression ast.Expression) value.Value {

	switch e := expression.(type) {
	case *ast.CommaExpressions:
		return c.compileCommaExpression(e)
	case *ast.PrefixOperation:
		return c.compilePrefixExpression(e)

	case *ast.AnonymousFunction:
		return c.compileAnonymousFunction(e)

	case *ast.Integer:
		{
			theType := c.ToLLVMType(e.Type)
			integerType := theType.(*types.IntType)
			return constant.NewInt(integerType, e.Value)
		}

	case *ast.Float:
		{
			theType := c.ToLLVMType(e.Type)
			floatType := theType.(*types.FloatType)
			return constant.NewFloat(floatType, e.Value)
		}

	case *ast.BinaryOperation:
		{
			return c.compileBinaryExpression(e)
		}

	case *ast.IndexAccess:
		{
			return c.compileIndexAccess(e)
		}

	case *ast.Call:
		{
			return c.compileFunctionCall(e)
		}

	case *ast.BuiltinCall:
		{
			return c.compileBuiltinFunctionCall(e)
		}

	case *ast.Identifier:
		{
			return c.compileIdentifier(e)
		}

	case *ast.StructLiteral:
		{
			return c.compileStructLiteral(e)
		}

	case *ast.ArrayLiteral:
		{
			return c.compileArrayLiteral(e)
		}

	case *ast.StringLiteral:
		return c.compileStringLiteral(e)
	}

	return nil
}

func (c *Compiler) compileAnonymousFunction(anonymousFunction *ast.AnonymousFunction) value.Value {
	name := "func." + random.RandomString(10)
	var namedValues []NamedValue
	for _, capturedVariableName := range anonymousFunction.CapturedVariables {
		variable := c.loadIfPointer(c.retrieveVariable(capturedVariableName))
		global := c.m.NewGlobalDef("func.captured."+capturedVariableName+random.RandomString(5), constant.NewUndef(variable.Type()))
		c.block().NewStore(variable, global)
		namedValues = append(namedValues, NamedValue{Name: capturedVariableName, Value: global})
	}

	anonymousFunction.FunctionType.ExternalName = name
	anonymousFunction.FunctionType.Name = name
	function := anonymousFunction
	c.compileFunctionDeclaration(name, function, namedValues...)
	funk := c.retrieveVariable(name)
	return funk
}

func (c *Compiler) compileStringLiteral(stringLiteral *ast.StringLiteral) value.Value {
	charArray := constant.NewCharArrayFromString(stringLiteral.Value + string(byte(0)))
	if c.currentFunction == nil {
		return charArray
	}
	globalDef := c.m.NewGlobalDef("string.literal."+random.RandomString(10), charArray)
	c.doNotLoadIntoMemory = true
	ptr := c.block().NewGetElementPtr(types.NewArray(uint64(len(stringLiteral.Value)+1), types.I8), globalDef, zero, zero)
	ptr.InBounds = true
	return ptr
}

func (c *Compiler) compileArrayLiteral(arrayLiteral *ast.ArrayLiteral) value.Value {
	arrayType := c.ToLLVMType(arrayLiteral.Type).(*types.ArrayType)
	allocaInstance := c.block().NewAlloca(arrayType)
	for index, value := range arrayLiteral.Values {
		loadedValue := c.loadIfPointer(c.compileExpression(value))
		integerIndex := constant.NewInt(types.I32, int64(index))
		address := c.block().NewGetElementPtr(arrayType, allocaInstance, zero, integerIndex)
		address.InBounds = true
		c.block().NewStore(loadedValue, address)
	}
	c.doNotAllocate = true
	return allocaInstance
}

func (c *Compiler) compilePrefixExpression(prefix *ast.PrefixOperation) value.Value {
	prefixValue := c.compileExpression(prefix.Right)
	if prefix.Operation == ops.Subtract {
		prefixValue = c.loadIfPointer(prefixValue)
		var negativeOne value.Value
		if ctypes.IsFloat(prefix.GetType()) {
			negativeOne = constant.NewFloat(prefixValue.Type().(*types.FloatType), -1)
			return c.block().NewFMul(prefixValue, negativeOne)
		}

		negativeOne = constant.NewInt(prefixValue.Type().(*types.IntType), -1)

		return c.block().NewMul(prefixValue, negativeOne)
	}

	if prefix.Operation == ops.AddOne || prefix.Operation == ops.SubtractOne {
		if !types.IsPointer(prefixValue.Type()) {
			c.exitErrorExpression("the right side of the value was not a variable to store on", prefix)
		}

		var newValue value.Value
		if prefix.Operation == ops.AddOne {
			// log.Println(c.doNotAllocate, c.doNotLoadIntoMemory, prefixValue, prefix)
			newValue = c.addOne(c.loadIfPointer(prefixValue))
		} else if prefix.Operation == ops.SubtractOne {
			newValue = c.subtractOne(c.loadIfPointer(prefixValue))
		}

		c.block().NewStore(newValue, prefixValue)
		return prefixValue
	}

	if prefix.Operation == ops.BinaryAND {
		// This means that we are referencing a variable that either:
		// * Has just been returned by a function, (func whatever() i32; whatever(); << we are on i32 instead of *i32)
		// * Or is just a literal, like: variable := &3 (i32 instead of *i32)
		//
		if !types.IsPointer(prefixValue.Type()) || c.doNotLoadIntoMemory {
			prefixValueTmp := c.block().NewAlloca(prefixValue.Type())
			c.block().NewStore(prefixValue, prefixValueTmp)
			prefixValue = prefixValueTmp
		}

		c.doNotLoadIntoMemory = false
		c.doNotAllocate = false
		allocatedValue := c.block().NewAlloca(prefixValue.Type())
		c.block().NewStore(prefixValue, allocatedValue)
		return allocatedValue
	}

	if prefix.Operation == ops.Multiply {
		prefixValue = c.loadIfPointer(prefixValue)
		return prefixValue
	}

	if prefix.Operation == ops.Bang {
		return c.block().NewICmp(enum.IPredEQ, c.toBool(c.loadIfPointer(prefixValue)), zero)
	}

	c.exitErrorExpression("unknown prefix", prefix)
	return nil
}

func (c *Compiler) exitInternalError(message string) {
	logger.Error("Internal error", "Unknown internal error has happened, check below for more details", "\n"+message)
	os.Exit(1)
}

func (c *Compiler) exit(message string) {
	logger.Error("Compiler error", "Unknown internal error has happened, check below for more details", "\n"+message)
	os.Exit(1)
}

func (c *Compiler) exitErrorExpression(message string, node ast.Expression) {
	logger.Error("Compiler", message, fmt.Sprintf("\non [%d:%d]: ", node.GetToken().Line, node.GetToken().Position)+node.String())
	os.Exit(1)
}

// NOTE: change of plans, we are now loading identifiers stack references and if the caller needs it we
// load it there
func (c *Compiler) compileIdentifier(id *ast.Identifier) value.Value {
	return c.retrieveVariable(id.Name)
}

func (c *Compiler) retrieveVariable(name string) value.Value {
	identifier := c.retrieveLocalVariable(name)
	if identifier != nil {
		if !types.IsPointer(identifier.Type()) {
			c.doNotLoadIntoMemory = true
			if types.IsArray(identifier.Type()) {
				global := identifier.(constant.Constant)
				array := c.block().NewAlloca(global.Type())
				c.block().NewStore(global, array)
				pointer := c.block().NewBitCast(array, types.NewPointer(global.Type().(*types.ArrayType).ElemType))
				lastAlloca := c.block().NewAlloca(pointer.Type())
				c.variables.Add(name, lastAlloca)
				c.block().NewStore(pointer, lastAlloca)
				c.doNotLoadIntoMemory = false
				return lastAlloca
			}
		}
		return identifier
	}

	if fn, ok := c.globalVariables[name]; ok {
		if !ctypes.IsFunction(fn.Type) {
			if existing, exists := c.globalVariables[name+".global"]; exists {
				return existing.Value
			}

			global := fn.Value.(constant.Constant)

			// If it is an array we must allocate it as a global and then cast it as a pointer
			// and allocate it again so we can use it normally, we cache that result into globalVariables + .global
			if types.IsArray(global.Type()) {
				array := c.m.NewGlobalDef(name+".global_alloca", global)
				pointer := c.block().NewBitCast(array, types.NewPointer(global.Type().(*types.ArrayType).ElemType))
				globalDefinition := c.m.NewGlobalDef(name+".global", constant.NewNull(pointer.Type().(*types.PointerType)))
				c.globalVariables[name+".global"] = &Value{Value: globalDefinition, Type: fn.Type}
				c.block().NewStore(pointer, globalDefinition)
				return globalDefinition
			}

			if fn.Constant {
				return global
			}

			globalDefPointer := c.m.NewGlobalDef(name+".global", global)
			globalDefPointer.Immutable = false
			c.globalVariables[name+".global"] = &Value{Value: globalDefPointer, Type: fn.Type}

			return globalDefPointer
		}

		c.doNotLoadIntoMemory = true
		return fn.Value
	}

	c.exitInternalError(fmt.Sprintf("Variable %s doesn't exist.", name))
	panic("")
}

func (c *Compiler) retrieveLocalVariable(name string) value.Value {
	return c.variables.Get(name)
}

/// Function calls
func (c *Compiler) compileBuiltinFunctionCall(ast *ast.BuiltinCall) value.Value {
	if fun, ok := c.builtins[ast.Name]; ok {
		return fun(c, ast)
	}

	c.exitInternalError("undefined builtin function @" + ast.Name)
	panic("")
}

func (c *Compiler) compileFunctionCall(ast *ast.Call) value.Value {
	left := c.compileExpression(ast.Left)
	funk := c.loadIfPointer(left)
	arguments := make([]value.Value, 0, len(ast.Parameters))
	for _, argument := range ast.Parameters {
		compiledValue := c.compileExpression(argument)
		loadedValue := c.loadIfPointer(compiledValue)
		arguments = append(arguments, loadedValue)
	}
	thing := c.block().NewCall(funk, arguments...)
	c.doNotLoadIntoMemory = true
	return thing
}

func (c *Compiler) allocateIfNotPointer(v value.Value) value.Value {
	t := v.Type()
	if _, isPtr := v.Type().(*types.PointerType); isPtr {
		return v
	}

	alloca := c.block().NewAlloca(t)
	c.block().NewStore(v, alloca)
	return alloca
}

func (c *Compiler) compileAssignment(assignment *ast.AssignmentStatement) {
	l := c.compileExpression(assignment.Left)
	// Reset state of doNotAllocate and doNotLoadIntoMemory so it doesn't affect second
	// compile expression
	c.doNotAllocate = false
	c.doNotLoadIntoMemory = false
	rightElement := c.compileExpression(assignment.Expression)

	if multipleValues, isMultipleValueAssignment := assignment.Left.(*ast.CommaExpressions); isMultipleValueAssignment {
		leftPointer := c.allocateIfNotPointer(l)
		rightPointer := c.allocateIfNotPointer(rightElement)
		c.doNotAllocate = false
		c.doNotLoadIntoMemory = false
		for i := range multipleValues.Expressions {
			rightValue := c.getStructField(rightPointer, i)
			leftAddress := c.getStructField(leftPointer, i)
			c.block().NewStore(
				rightValue,
				c.bitcastIfUnion(multipleValues.Expressions[i].GetType(), leftAddress, types.NewPointer(rightValue.Type())),
			)
		}
		return
	}

	r := c.loadIfPointer(rightElement)
	c.block().NewStore(r, c.bitcastIfUnion(assignment.Left.GetType(), l, types.NewPointer(r.Type())))
}

func (c *Compiler) compileIndexAccess(access *ast.IndexAccess) value.Value {
	leftArray := c.compileExpression(access.Left)
	// keep do not load into memory before compiling expression
	doNotLoadLeftArray := c.doNotLoadIntoMemory
	index := c.loadIfPointer(c.compileExpression(access.Access))

	// If it's an array do not load into memory, just calculate offset
	if types.IsPointer(leftArray.Type()) && types.IsArray(leftArray.Type().(*types.PointerType).ElemType) {
		// zero as first offset because we are calculating first pointer, then index
		element := c.block().NewGetElementPtr(leftArray.Type().(*types.PointerType).ElemType, leftArray, zero, index)
		element.InBounds = true
		return element
	}

	c.doNotLoadIntoMemory = doNotLoadLeftArray
	leftArray = c.loadIfPointer(leftArray)
	return c.calculatePointerOffset(leftArray, index)
}

func (c *Compiler) calculatePointerOffset(pointer value.Value, offset value.Value) value.Value {
	calculatedPointer := c.block().NewGetElementPtr(pointer.Type().(*types.PointerType).ElemType, pointer, offset)
	calculatedPointer.InBounds = true
	return calculatedPointer
}

func (c *Compiler) getStructField(pointerStruct value.Value, fieldPosition int) value.Value {
	pointerType := pointerStruct.Type().(*types.PointerType)
	ptr := c.block().NewGetElementPtr(
		pointerType.ElemType, pointerStruct, zero, constant.NewInt(types.I32, int64(fieldPosition)),
	)
	ptr.InBounds = true
	return c.loadIfPointer(ptr)
}

func unwrapArrayGepType(arr *types.ArrayType) types.Type {
	return types.NewPointer(arr.ElemType)
}

func (c *Compiler) loadIfPointer(val value.Value) value.Value {
	if c.doNotLoadIntoMemory {
		c.doNotLoadIntoMemory = false
		c.doNotAllocate = false
		return val
	}
	if types.IsPointer(val.Type()) {
		load := c.block().NewLoad(val.Type().(*types.PointerType).ElemType, val)
		// if fl, ok := val.Type().(*types.PointerType).ElemType.(*types.FloatType); ok {
		// 	if fl.Kind == types.FloatKindFloat {
		// 		load.Align = 4
		// 	} else {
		// 		load.Align = 8
		// 	}
		// } else if number, ok := val.Type().(*types.PointerType).ElemType.(*types.IntType); ok {
		// 	load.Align = ir.Align(number.BitSize / 8)
		// }
		return load
	}
	return val
}

func (c *Compiler) loadIfPointerWithAlignment(val value.Value, alignment int) value.Value {
	if c.doNotLoadIntoMemory {
		c.doNotLoadIntoMemory = false
		c.doNotAllocate = false
		return val
	}
	if types.IsPointer(val.Type()) {
		l := c.block().NewLoad(val.Type().(*types.PointerType).ElemType, val)
		l.Align = ir.Align(alignment)
		return l
	}
	return val
}

func (c *Compiler) compileStructLiteral(strukt *ast.StructLiteral) value.Value {
	module := c
	if strukt.Module != "" {
		module = c.modules[strukt.Module]
	}
	possibleStruct := module.types[strukt.Name]
	struktType, ok := possibleStruct.candiceType.(*ctypes.Struct)

	if !ok {
		c.exitInternalError(fmt.Sprintf("expected struct but got a %s", possibleStruct.candiceType))
		panic("")
	}

	// Allocate in stack memory
	struktValue := c.block().NewAlloca(possibleStruct.llvmType)

	for _, decl := range strukt.Values {
		// Get field position
		i, field := struktType.GetField(decl.Name)

		if anonymous, ok := field.(*ctypes.Anonymous); ok {
			field = c.types[anonymous.Name].candiceType
		}

		// Compile the expression to have the value
		compiledValue := c.compileExpression(decl.Expression)

		// Get the pointer pointing to the memory where we need to store in
		ptr := c.block().NewGetElementPtr(possibleStruct.llvmType, struktValue, zero, constant.NewInt(types.I32, int64(i)))
		ptr.InBounds = true

		// Unwrap value pointer
		compiledValue = c.loadIfPointer(c.bitcastIfUnion(field, compiledValue, ptr.Type()))

		// Store in the pointer the compiler value
		c.block().NewStore(compiledValue, ptr)
	}

	// Do not allocate on declaration/assignments because we are already allocating above.
	c.doNotAllocate = true

	return struktValue
}

/// Simple binary compilations
/// Making redundant and easy to understand globalVariables is better
/// than storing callbacks on a hashmap. Let's keep it simple.
func (c *Compiler) compileBinaryExpression(expr *ast.BinaryOperation) value.Value {
	switch expr.Operation {
	case ops.Modulo:
		{
			return c.compileModulo(expr)
		}

	case ops.Multiply:
		{
			return c.compileMultiply(expr)
		}

	case ops.Add:
		{
			return c.compileAdd(expr)
		}

	case ops.Divide:
		{
			return c.compileDivide(expr)
		}

	case ops.Subtract:
		{
			return c.compileSubtract(expr)
		}

	case ops.BinaryXOR:
		{
			return c.compileXorBinary(expr)
		}

	case ops.BinaryAND:
		{
			return c.compileAndBinary(expr)
		}

	case ops.BinaryOR:
		{
			return c.compileOrBinary(expr)
		}

	case ops.Dot:
		{
			return c.compileStructAccess(expr)
		}

	case ops.AND:
		{
			return c.compileAnd(expr)
		}

	case ops.OR:
		{
			return c.compileOr(expr)
		}

	case ops.LeftShift:
		{
			return c.compileShiftLeftBinary(expr)
		}

	case ops.RightShift:
		{
			return c.compileShiftRightBinary(expr)
		}

	default:
		return c.handleComparisonOperations(expr)
	}
}

func (c *Compiler) compileOr(and *ast.BinaryOperation) value.Value {
	returnValue := constant.NewBool(false)
	allocatedValue := c.block().NewAlloca(types.I1)
	c.block().NewStore(returnValue, allocatedValue)
	orBlock := c.currentFunction.NewBlock("or." + random.RandomString(10))
	truthBlock := c.currentFunction.NewBlock("truth_or." + random.RandomString(10))
	leaveBlock := c.currentFunction.NewBlock("leaveor." + random.RandomString(10))
	c.block().NewBr(orBlock)
	c.pushBlock(orBlock)
	leftExpression := c.loadIfPointer(c.compileExpression(and.Left))
	boolean := c.toBool(leftExpression)
	c.block().NewStore(boolean, allocatedValue)
	c.block().NewCondBr(boolean, leaveBlock, truthBlock)
	orBlock = c.popBlock()
	c.pushBlock(truthBlock)
	value := c.toBool(c.loadIfPointer(c.compileExpression(and.Right)))
	c.block().NewStore(value, allocatedValue)
	c.block().NewBr(leaveBlock)
	truthBlock = c.popBlock()
	c.blocks[len(c.blocks)-1] = leaveBlock
	return allocatedValue
}

func (c *Compiler) compileAnd(and *ast.BinaryOperation) value.Value {
	returnValue := constant.NewBool(false)
	allocatedValue := c.block().NewAlloca(types.I1)
	c.block().NewStore(returnValue, allocatedValue)
	andBlock := c.currentFunction.NewBlock("and." + random.RandomString(10))
	truthBlock := c.currentFunction.NewBlock("truth_and." + random.RandomString(10))
	leaveBlock := c.currentFunction.NewBlock("leaveand." + random.RandomString(10))
	c.block().NewBr(andBlock)
	c.pushBlock(andBlock)
	leftExpression := c.loadIfPointer(c.compileExpression(and.Left))
	boolean := c.toBool(leftExpression)
	c.block().NewCondBr(boolean, truthBlock, leaveBlock)
	c.popBlock()
	c.pushBlock(truthBlock)
	value := c.toBool(c.loadIfPointer(c.compileExpression(and.Right)))
	c.block().NewStore(value, allocatedValue)
	c.block().NewBr(leaveBlock)
	c.popBlock()
	c.blocks[len(c.blocks)-1] = leaveBlock
	return allocatedValue
}

func getName(expr ast.Expression) (string, bool) {
	if bin, ok := expr.(*ast.BinaryOperation); ok {
		return bin.Left.(*ast.Identifier).Name, false
	}
	if bin, ok := expr.(*ast.Identifier); ok {
		return bin.Name, true
	}
	panic("INTERNAL ERROR: " + expr.String())
}

func (c *Compiler) compileModuleAccess(expr *ast.BinaryOperation) value.Value {
	moduleName := expr.Left.(*ast.Identifier).Name
	module := c.modules[moduleName]
	// In case compileIdentifier tries to generate code or local variables
	prevStack, prevBlock := module.variables, module.blocks
	module.variables = c.variables
	module.blocks = c.blocks
	identifier := module.compileIdentifier(expr.Right.(*ast.Identifier))
	module.variables = prevStack
	module.blocks = prevBlock
	c.doNotLoadIntoMemory = module.doNotLoadIntoMemory
	module.doNotLoadIntoMemory = false
	return identifier
}

func (c *Compiler) compileStructAccess(expr *ast.BinaryOperation) value.Value {
	if _, isModule := expr.Left.GetType().(*semantic.Semantic); isModule {
		return c.compileModuleAccess(expr)
	}

	leftStructPossibleNonPtr := c.compileExpression(expr.Left)
	currentCandiceType := expr.Left.GetType()
	var candiceType ctypes.FieldType
	leftStruct := leftStructPossibleNonPtr

	if !types.IsPointer(leftStructPossibleNonPtr.Type()) {
		leftStruct = c.block().NewAlloca(leftStructPossibleNonPtr.Type())
		c.block().NewStore(leftStructPossibleNonPtr, leftStruct)
	}

	if s, ok := leftStruct.Type().(*types.PointerType); ok {
		// Load if it's a pointer
		if types.IsPointer(s.ElemType) {
			leftStruct = c.loadIfPointer(leftStruct)
			s, ok = leftStruct.Type().(*types.PointerType)
			if !ok {
				c.exitInternalError("not a struct " + leftStruct.Type().String() + " " + expr.String())
			}
		}

		if ctypes.IsPointer(currentCandiceType) {
			currentCandiceType = currentCandiceType.(*ctypes.Pointer).Inner
		}
		if anonymous, ok := currentCandiceType.(*ctypes.Anonymous); ok {
			if anonymous.Modules != nil && len(anonymous.Modules) != 0 {
				module := anonymous.Modules[0]
				candiceType = c.modules[module].types[s.ElemType.Name()].candiceType.(ctypes.FieldType)
			} else {
				candiceType = c.types[s.ElemType.Name()].candiceType.(ctypes.FieldType)
			}
		} else {
			candiceType = currentCandiceType.(ctypes.FieldType)
		}
	}

	rightName, _ := getName(expr.Right)
	i, field := candiceType.GetField(ast.RetrieveID(rightName))
	if _, isStruct := candiceType.(*ctypes.Struct); isStruct {
		// This is always a pointer
		inner := leftStruct.Type().(*types.PointerType).ElemType
		// Zero to calculate address of pointer and then calculate the address for the field
		ptr := c.block().NewGetElementPtr(inner, leftStruct, zero, constant.NewInt(types.I32, int64(i)))
		ptr.InBounds = true
		leftStruct = ptr
	} else {
		// we know this is a union
		leftStruct = c.block().NewBitCast(leftStruct, types.NewPointer(c.ToLLVMType(field)))
	}

	// We are possibly not doing c.loadIfPointer calls on this function.
	// But we still need to reset these flags.
	c.doNotAllocate = false
	c.doNotLoadIntoMemory = false

	return leftStruct
}

// compileAdd can handle *i8 sums, memory offset accessing and numeric operations
// numeric operations should have the same type, else it will fail
func (c *Compiler) compileAdd(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))

	// String concatenation
	if types.IsPointer(rightValue.Type()) {
		c.doNotLoadIntoMemory = true
		return c.concatenateMemoryI8(leftValue, rightValue)
	}

	// Memory access via '+'
	if types.IsPointer(leftValue.Type()) {
		newPointer := c.calculatePointerOffset(leftValue, rightValue)
		// return this pointer on another stack register because it's going to get used
		toReturnPointer := c.block().NewAlloca(newPointer.Type())
		c.block().NewStore(newPointer, toReturnPointer)
		return toReturnPointer
	}

	// numeric operations
	if types.IsFloat(leftValue.Type()) {
		return c.block().NewFAdd(leftValue, rightValue)
	}

	return c.block().NewAdd(leftValue, rightValue)
}

func (c *Compiler) addOne(v value.Value) value.Value {
	if types.IsFloat(v.Type()) {
		return c.block().NewFAdd(v, constant.NewFloat(v.Type().(*types.FloatType), 1.0))
	}

	return c.block().NewAdd(v, constant.NewInt(v.Type().(*types.IntType), 1))
}

func (c *Compiler) subtractOne(v value.Value) value.Value {
	if types.IsFloat(v.Type()) {
		return c.block().NewFSub(v, constant.NewFloat(v.Type().(*types.FloatType), 1.0))
	}
	return c.block().NewSub(v, constant.NewInt(v.Type().(*types.IntType), 1))
}

func (c *Compiler) compileMultiply(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	if types.IsFloat(leftValue.Type()) {
		return c.block().NewFMul(leftValue, rightValue)
	}
	return c.block().NewMul(leftValue, rightValue)
}

func (c *Compiler) compileSubtract(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	if types.IsPointer(leftValue.Type()) {
		rightValue = c.block().NewMul(rightValue, constant.NewInt(rightValue.Type().(*types.IntType), -1))
		newPointer := c.calculatePointerOffset(leftValue, rightValue)
		// return this pointer on another stack register because it's probably going to get used
		toReturnPointer := c.block().NewAlloca(newPointer.Type())
		c.block().NewStore(newPointer, toReturnPointer)
		return toReturnPointer
	}

	if types.IsFloat(leftValue.Type()) {
		return c.block().NewFSub(leftValue, rightValue)
	}
	return c.block().NewSub(leftValue, rightValue)
}

func (c *Compiler) compileDivide(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	if types.IsInt(leftValue.Type()) {
		if _, isUnsigned := expr.Type.(*ctypes.UInteger); isUnsigned {
			return c.block().NewUDiv(leftValue, rightValue)
		}
		return c.block().NewSDiv(leftValue, rightValue)
	}
	if types.IsFloat(leftValue.Type()) {
		return c.block().NewFDiv(leftValue, rightValue)
	}
	c.exitErrorExpression("can't divide these types", expr)
	panic("")
}

func (c *Compiler) compileModulo(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	if types.IsInt(leftValue.Type()) {
		if _, isUnsigned := expr.Type.(*ctypes.UInteger); isUnsigned {
			return c.block().NewURem(leftValue, rightValue)
		}
		return c.block().NewSRem(leftValue, rightValue)
	}
	if types.IsFloat(leftValue.Type()) {
		return c.block().NewFRem(leftValue, rightValue)
	}
	c.exitErrorExpression("can't divide these types", expr)
	panic("")
}

func (c *Compiler) compileAndBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewAnd(leftValue, rightValue)
}

func (c *Compiler) compileOrBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewOr(leftValue, rightValue)
}

func (c *Compiler) compileXorBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewXor(leftValue, rightValue)
}

func (c *Compiler) compileShiftRightBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewLShr(leftValue, rightValue)
}

func (c *Compiler) compileShiftLeftBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewShl(leftValue, rightValue)
}

func (c *Compiler) handleCast(call *ast.BuiltinCall) value.Value {
	typeParameter := call.TypeParameters[0]
	toReturnType := c.ToLLVMType(typeParameter)
	variable := c.compileExpression(call.Parameters[0])

	// TODO: Array to integer
	if types.IsPointer(variable.Type()) && types.IsArray(variable.Type().(*types.PointerType).ElemType) {
		// We don't want people to take this pointer and load it into memory
		// Test this because I'm not sure if this works!
		c.doNotLoadIntoMemory = true
		// We are not running c.loadIfPointer so put this to false to reset it
		c.doNotAllocate = false
		ptr := c.block().NewGetElementPtr(variable.Type().(*types.PointerType).ElemType, variable, zero, zero)
		bc := c.block().NewBitCast(ptr, toReturnType)
		return bc
	}
	variable = c.loadIfPointer(variable)

	if ctypes.IsNumeric(call.TypeParameters[0]) && ctypes.IsNumeric(call.Parameters[0].GetType()) {
		if ctypes.IsFloat(call.TypeParameters[0]) != ctypes.IsFloat(call.Parameters[0].GetType()) {
			return c.handleFloatIntCast(call.TypeParameters[0], call.Parameters[0].GetType(), variable, toReturnType)
		}
		return c.handleNumericBitCast(toReturnType, variable)
	}

	if ctypes.IsNumeric(call.TypeParameters[0]) && ctypes.IsPointer(call.Parameters[0].GetType()) {
		integer := c.block().NewPtrToInt(variable, toReturnType)
		if _, isFloat := call.TypeParameters[0].(*ctypes.Float); isFloat {
			return c.block().NewSIToFP(integer, toReturnType)
		}
		return integer
	}

	if ctypes.IsPointer(call.TypeParameters[0]) && ctypes.IsNumeric(call.Parameters[0].GetType()) { ///types.IsInt(variable.Type()) {
		// do not load into memory because our conversion doesn't let us put the pointer above its level
		// for example we should target to return i32** instead of i32* (if our pointer type was i32*)
		// if we had this to false
		c.doNotLoadIntoMemory = true
		if fl, isFloat := call.Parameters[0].GetType().(*ctypes.Float); isFloat {
			variable = c.block().NewFPToSI(variable, types.NewInt(uint64(fl.BitSize)))
		}
		value := c.block().NewIntToPtr(variable, toReturnType)
		return value
	}

	if ctypes.IsPointer(call.TypeParameters[0]) && types.IsPointer(variable.Type()) {
		c.doNotLoadIntoMemory = true
		c.doNotAllocate = false
		return c.block().NewBitCast(variable, toReturnType)
	}

	c.exitInternalError("cant convert yet to this " + call.String() + "\n" + call.Parameters[0].GetType().String() + " " + variable.Type().String())
	panic("")
}

func (c *Compiler) compileSwitchStatement(switchStatement *ast.SwitchStatement) {
	condition := c.loadIfPointer(c.compileExpression(switchStatement.Condition))
	var strandedBlocks []*ir.Block
	leaveBlock := ir.NewBlock("leaveSwitchBlock." + random.RandomString(10))
	var defaultBlock *ir.Block
	if switchStatement.Default != nil {
		defaultBlock = c.currentFunction.NewBlock("default." + random.RandomString(10))
		strandedBlocks = append(strandedBlocks, c.compileBlock(switchStatement.Default, defaultBlock))
	} else {
		defaultBlock = leaveBlock
	}

	var cases []*ir.Case

	casesId := random.RandomString(10)
	for i, caseStatement := range switchStatement.Cases {
		possibleNonConstantExpression := c.compileConstantExpression(caseStatement.Case)
		if constant, isConstant := possibleNonConstantExpression.(constant.Constant); isConstant {
			caseBlock := c.currentFunction.NewBlock(fmt.Sprintf("case-%d-%s", i, casesId))
			irCase := ir.NewCase(constant, caseBlock)
			cases = append(cases, irCase)
			strandedBlocks = append(strandedBlocks, c.compileBlock(caseStatement.Block, caseBlock))
		} else {
			logger.Warning("You are using an experimental part of Candice, you can't use non-constant expressions right now on switch cases\n")
			c.exit("Exiting compiler because of non-constant expression")
		}
	}

	c.block().Term = c.block().NewSwitch(condition, defaultBlock, cases...)

	for _, strandedBlock := range strandedBlocks {
		if strandedBlock.Term == nil {
			strandedBlock.NewBr(leaveBlock)
		}
	}

	c.blocks[len(c.blocks)-1] = leaveBlock
	c.currentFunction.Blocks = append(c.currentFunction.Blocks, leaveBlock)
}

func (c *Compiler) compileCommaExpression(commaExpression *ast.CommaExpressions) value.Value {
	var values []value.Value
	for _, expression := range commaExpression.Expressions {
		v := c.compileExpression(expression)
		if !commaExpression.IsAssignment {
			v = c.loadIfPointer(v)
		} else {
			c.doNotAllocate = false
			c.doNotLoadIntoMemory = false
		}

		values = append(values, v)
	}

	var ptr *types.PointerType
	if commaExpression.IsAssignment {
		var typeList []types.Type
		for _, value := range values {
			typeList = append(typeList, value.Type())
		}
		strukt := types.NewStruct(typeList...)
		ptr = types.NewPointer(strukt)
	} else {
		ptr = types.NewPointer(c.ToLLVMType(commaExpression.Type))
	}

	strukt := c.block().NewAlloca(ptr.ElemType)
	for i := range commaExpression.Expressions {
		address := c.block().NewGetElementPtr(strukt.ElemType, strukt, zero, constant.NewInt(types.I32, int64(i)))
		c.block().NewStore(values[i], address)
	}
	// c.doNotLoadIntoMemory = true
	return strukt
}

func (c *Compiler) compileMultipleDeclarationStatement(m *ast.MultipleDeclarationStatement) {
	strukt := c.compileExpression(m.Expression)
	c.doNotLoadIntoMemory = false
	c.doNotAllocate = false
	if !types.IsPointer(strukt.Type()) {
		struktPtr := c.block().NewAlloca(strukt.Type())
		c.block().NewStore(strukt, struktPtr)
		strukt = struktPtr
	}

	for i := range m.Names {
		address :=
			c.block().NewGetElementPtr(
				strukt.Type().(*types.PointerType).ElemType, strukt, zero, constant.NewInt(types.I32, int64(i)),
			)
		c.declare(m.Names[i], address)
	}
}
