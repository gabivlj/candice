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

	builtins        map[string]func(*Compiler, *ast.BuiltinCall) value.Value
	globalVariables map[string]*Value

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

	free := c.m.NewFunc("free", types.Void, ir.NewParam("", types.I8Ptr))
	c.globalBuiltinDefinitions["free"] = free
	free.Sig.Variadic = true
	free.CallingConv = enum.CallingConvC
	c.builtins["free"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		ptr := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		returnedValue := c.block().NewCall(free, c.block().NewBitCast(ptr, types.I8Ptr))
		return returnedValue
	}

	realloc := c.m.NewFunc("realloc", types.I8Ptr, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I64))
	c.globalBuiltinDefinitions["realloc"] = realloc
	realloc.Sig.Variadic = true
	realloc.CallingConv = enum.CallingConvC

	c.builtins["realloc"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		typeParameter := call.GetType()
		toReturnType := c.ToLLVMType(typeParameter)
		ptr := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		length := c.loadIfPointer(c.compileExpression(call.Parameters[1]))
		length = c.handleIntegerCast(types.I64, length)
		totalSize := c.block().NewMul(length, constant.NewInt(types.I64, typeParameter.SizeOf()))
		returnedValue := c.block().NewCall(realloc, c.block().NewBitCast(ptr, types.I8Ptr), totalSize)
		castedValue := c.block().NewBitCast(returnedValue, toReturnType)
		alloca := c.block().NewAlloca(castedValue.Type())
		c.block().NewStore(castedValue, alloca)
		return alloca
	}

	c.builtins["sizeof"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		return constant.NewInt(types.I32, call.TypeParameters[0].SizeOf())
	}

	printf := c.m.NewFunc(
		"printf",
		types.I32,
		ir.NewParam("", types.NewPointer(types.I8)),
	)
	c.globalBuiltinDefinitions["printf"] = printf
	printf.Sig.Variadic = true
	printf.CallingConv = enum.CallingConvC
	c.builtins["print"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		expressions := make([]value.Value, len(call.Parameters)+1)
		for i := 0; i < len(call.Parameters); i++ {
			expressions[i+1] = c.loadIfPointer(c.compileExpression(call.Parameters[i]))
		}
		constantString := strings.Builder{}
		for i := range call.Parameters {
			t := expressions[i+1].Type()
			constantString.WriteString(c.getFormatString(expressions, t, call, i))
		}
		s := constantString.String()
		expressions[0] = c.createString(s)
		c.block().NewCall(printf, expressions...)
		return constant.NewUndef(types.Void)
	}

	malloc := c.m.NewFunc(
		"malloc",
		types.NewPointer(types.I8),
		ir.NewParam("", types.I64),
	)
	c.globalBuiltinDefinitions["malloc"] = malloc
	printf.CallingConv = enum.CallingConvC
	// alloc accepts one type parameter, and how many you want to allocate
	c.builtins["alloc"] = func(c *Compiler, call *ast.BuiltinCall) value.Value {
		typeParameter := call.TypeParameters[0]
		toReturnType := types.NewPointer(c.ToLLVMType(typeParameter))
		length := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		length = c.handleIntegerCast(types.I64, length)
		totalSize := c.block().NewMul(length, constant.NewInt(types.I64, typeParameter.SizeOf()))
		returnedValue := c.block().NewCall(malloc, totalSize)
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
	logger.Success("Finished compiling on Candice.")
	fd, _ := os.Create(".intermediate_output.ll")
	_, _ = c.m.WriteTo(fd)
	endFlags := []string{".intermediate_output.ll", "-o", output}
	endFlags = append(endFlags, flags...)
	cmd := exec.Command(cxx, endFlags...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	logger.Success("Compiling on " + cxx)
	err := cmd.Run()
	logger.Success("Compiled on " + cxx)
	if err != nil {
		return errors.New("error compiling with clang:\n" + stdout.String())
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
			localCompiler.Compile(module.Root)
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

	if f, ok := c.globalVariables[funcType.ExternalName]; ok {
		c.globalVariables[funcType.Name] = f
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
	c.globalVariables[funcType.Name] = funk
	c.globalVariables[funcType.ExternalName] = funk
}

func (c *Compiler) compileReturn(ret *ast.ReturnStatement) {
	if ret.Expression == nil {
		c.block().NewRet(nil)
		return
	}

	toReturn := c.compileExpression(ret.Expression)
	c.block().NewRet(c.loadIfPointer(toReturn))
}

// defines function type without compiling its body
func (c *Compiler) compileFunctionType(name string, funk *ast.FunctionDeclarationStatement) *ir.Func {
	if fn, ok := c.globalVariables[name]; ok {
		return fn.Value.(*ir.Func)
	}

	// Declare params LLVM IR
	params := make([]*ir.Param, 0, len(funk.FunctionType.Parameters))
	for i, param := range funk.FunctionType.Parameters {
		t := c.ToLLVMType(param)
		name := funk.FunctionType.Names[i]
		params = append(params, ir.NewParam(name, t))
	}

	if funk.FunctionType.IsMainFunction() && funk.FunctionType.Return == ctypes.VoidType || funk.FunctionType.Return == nil {
		funk.FunctionType.Return = ctypes.I32
	}

	// Declare llvmFunction
	llvmFunction := c.m.NewFunc(funk.FunctionType.Name, c.ToLLVMType(funk.FunctionType.Return), params...)

	if funk.FunctionType.RedefineWithOriginalName {
		llvmFunctionExtern := c.m.NewFunc(funk.FunctionType.ExternalName, c.ToLLVMType(funk.FunctionType.Return), params...)
		llvmFunctionExtern.CallingConv = enum.CallingConvC
		c.globalVariables[funk.FunctionType.ExternalName] = &Value{
			Value: llvmFunctionExtern,
			Type:  funk.FunctionType,
		}
	}

	// Create function
	c.globalVariables[name] = &Value{
		Value: llvmFunction,
		Type:  funk.FunctionType,
	}

	return llvmFunction
}

// compiles entire function
func (c *Compiler) compileFunctionDeclaration(name string, funk *ast.FunctionDeclarationStatement) {
	llvmFunction := c.compileFunctionType(name, funk)

	// Create a main block to the function
	c.pushBlock(llvmFunction.NewBlock(funk.FunctionType.Name))

	// Create a variable stack isolated from the rest of variable definitions
	// c.stacks = append(c.stacks, map[string]value.Value{})

	// Declare parameters IR
	for _, param := range llvmFunction.Params {
		allocatedParameter := c.block().NewAlloca(param.Type())
		c.block().NewStore(param, allocatedParameter)
		c.declare(param.Name(), allocatedParameter)
	}

	// Set it as current function
	prevFunction := c.currentFunction
	c.currentFunction = llvmFunction

	// Compile block
	for _, statement := range funk.Block.Statements {
		c.Compile(statement)
	}

	lastBlock := c.currentFunction.Blocks[len(c.currentFunction.Blocks)-1]
	// If return hasn't been declared, declare a void return
	if lastBlock.Term == nil {
		if funk.FunctionType.IsMainFunction() {
			lastBlock.NewRet(zero)
		} else {
			lastBlock.NewRet(nil)
		}
	}

	if funk.FunctionType.RedefineWithOriginalName {
		c.compileFunctionRedeclaration(c.currentFunction.Blocks, funk)
	}

	// Pop block, stack and restore current function
	c.popBlock()
	c.currentFunction = prevFunction
}

func (c *Compiler) compileFunctionRedeclaration(otherFunctionBlocks []*ir.Block, funk *ast.FunctionDeclarationStatement) {
	llvmFunction := c.retrieveVariable(funk.FunctionType.ExternalName).(*ir.Func)
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
		currentBlock = c.compileBlock(elseIf.Block, currentBlock)
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
	lastJumpToCondition.NewBr(leaveBlock)

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
	condition.NewCondBr(valueCondition, mainLoop, leave)
	c.popBlock()

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
	valueCompiled := c.compileExpression(decl.Expression)
	if c.currentFunction == nil {
		if _, isConstant := valueCompiled.(constant.Constant); isConstant {
			c.globalVariables[decl.Name] = &Value{Value: valueCompiled, Type: decl.Type}
			// c.globalVariables[decl.Name+".global"] = &Value{Value: c.m.NewGlobalDef(decl.Name, constant.NewNull(types.NewPointer(valueCompiled.Type()))), Type: decl.Type}
		} else {
			fmt.Println("Candice does not support constant expressions and operations yet.")
			os.Exit(1)
		}

		return
	}
	var val value.Value
	if !c.doNotAllocate {
		val = c.block().NewAlloca(t)
		c.block().NewStore(c.loadIfPointer(valueCompiled), val)
	} else {
		c.doNotAllocate = false
		val = valueCompiled
	}

	c.declare(decl.Name, val)
}

func (c *Compiler) declare(name string, value value.Value) {
	c.variables.Add(name, value)
}

func (c *Compiler) compileStruct(strukt *ast.StructStatement) {
	c.compileType(strukt.Type.Name, strukt.Type)
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
	anonymousFunction.FunctionType.ExternalName = name
	anonymousFunction.FunctionType.Name = name
	function := (*ast.FunctionDeclarationStatement)(anonymousFunction)
	c.compileFunctionDeclaration(name, function)
	funk := c.retrieveVariable(name)
	return funk
}

func (c *Compiler) compileStringLiteral(stringLiteral *ast.StringLiteral) value.Value {
	charArray := constant.NewCharArrayFromString(stringLiteral.Value + string(byte(0)))
	globalDef := c.m.NewGlobalDef("string.literal."+random.RandomString(10), charArray)
	if c.currentFunction == nil {
		return charArray
	}
	c.doNotLoadIntoMemory = true
	return c.block().NewGetElementPtr(types.NewArray(uint64(len(stringLiteral.Value)+1), types.I8), globalDef, zero, zero)
}

func (c *Compiler) compileArrayLiteral(arrayLiteral *ast.ArrayLiteral) value.Value {
	arrayType := c.ToLLVMType(arrayLiteral.Type).(*types.ArrayType)
	allocaInstance := c.block().NewAlloca(arrayType)
	for index, value := range arrayLiteral.Values {
		loadedValue := c.loadIfPointer(c.compileExpression(value))
		integerIndex := constant.NewInt(types.I32, int64(index))
		address := c.block().NewGetElementPtr(arrayType, allocaInstance, zero, integerIndex)
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
		// * Or is just a literal, like: variable := &3 << i32 instead of *i32
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
		prefixValue = c.block().NewLoad(prefixValue.Type().(*types.PointerType).ElemType, prefixValue)
		return prefixValue
	}

	if prefix.Operation == ops.Bang {
		return c.block().NewICmp(enum.IPredEQ, c.toBool(c.loadIfPointer(prefixValue)), zero)
	}

	c.exitErrorExpression("unknown prefix", prefix)
	return nil
}

func (c *Compiler) exit(message string) {
	logger.Error("Internal error", "Unknown internal error has happened, check below for more details", "\n"+message)
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

func (c *Compiler) compileIdentifierReference(id *ast.Identifier) value.Value {
	return c.retrieveLocalVariable(id.Name)
}

func (c *Compiler) retrieveVariable(name string) value.Value {
	identifier := c.retrieveLocalVariable(name)
	if identifier != nil {
		return identifier
	}

	if fn, ok := c.globalVariables[name]; ok {
		if !ctypes.IsFunction(fn.Type) {
			if existing, exists := c.globalVariables[name+".global"]; exists {
				return existing.Value
			}

			global := fn.Value.(constant.Constant)

			// If it is an array we must allocate it as a global and then cast it as a pointer
			// so we can mutate the address
			if types.IsArray(global.Type()) {
				array := c.m.NewGlobalDef(name+".global_alloca", global)
				pointer := c.block().NewBitCast(array, types.NewPointer(global.Type().(*types.ArrayType).ElemType))
				globalDefinition := c.m.NewGlobalDef(name+".global", constant.NewNull(pointer.Type().(*types.PointerType)))
				c.globalVariables[name+".global"] = &Value{Value: globalDefinition, Type: fn.Type}
				c.block().NewStore(pointer, globalDefinition)
				return globalDefinition
			}

			globalDefPointer := c.m.NewGlobalDef(name+".global", global)
			globalDefPointer.Immutable = false
			c.globalVariables[name+".global"] = &Value{Value: globalDefPointer, Type: fn.Type}

			return globalDefPointer
		}

		c.doNotLoadIntoMemory = true
		return fn.Value
	}

	c.exit(fmt.Sprintf("Variable %s doesn't exist.", name))
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

	c.exit("undefined builtin function @" + ast.Name)
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

func (c *Compiler) compileAssignment(assignment *ast.AssignmentStatement) {
	l := c.compileExpression(assignment.Left)
	r := c.loadIfPointer(c.compileExpression(assignment.Expression))
	c.block().NewStore(r, l)
}

func (c *Compiler) compileIndexAccess(access *ast.IndexAccess) value.Value {
	leftArray := c.compileExpression(access.Left)
	doNotLoadLeftArray := c.doNotLoadIntoMemory
	index := c.loadIfPointer(c.compileExpression(access.Access))

	if types.IsPointer(leftArray.Type()) && types.IsArray(leftArray.Type().(*types.PointerType).ElemType) {
		return c.block().NewGetElementPtr(leftArray.Type().(*types.PointerType).ElemType, leftArray, zero, index)
	}

	c.doNotLoadIntoMemory = doNotLoadLeftArray

	leftArray = c.loadIfPointer(leftArray)
	pointer := c.block().NewGetElementPtr(leftArray.Type().(*types.PointerType).ElemType, leftArray, index)
	return pointer
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
		return c.block().NewLoad(val.Type().(*types.PointerType).ElemType, val)
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
		c.exit(fmt.Sprintf("expected struct but got a %s", possibleStruct.candiceType))
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
		var ptr value.Value = c.block().NewGetElementPtr(possibleStruct.llvmType, struktValue, zero, constant.NewInt(types.I32, int64(i)))

		// Unwrap value pointer
		compiledValue = c.loadIfPointer(compiledValue)

		// Store in the pointer the compiler value
		c.block().NewStore(compiledValue, ptr)
	}

	return struktValue
}

/// Simple binary compilations
/// Making redundant and easy to understand globalVariables is better
/// than storing callbacks on a hashmap. Let's keep it simple.
func (c *Compiler) compileBinaryExpression(expr *ast.BinaryOperation) value.Value {
	switch expr.Operation {
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
	andBlock = c.popBlock()
	c.pushBlock(truthBlock)
	value := c.toBool(c.loadIfPointer(c.compileExpression(and.Right)))
	c.block().NewStore(value, allocatedValue)
	c.block().NewBr(leaveBlock)
	truthBlock = c.popBlock()
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

	leftStruct := c.compileExpression(expr.Left)
	currentCandiceType := expr.Left.GetType()
	var candiceType *ctypes.Struct
	if s, ok := leftStruct.Type().(*types.PointerType); ok {
		if types.IsPointer(s.ElemType) {
			leftStruct = c.loadIfPointer(leftStruct)
			s, ok = leftStruct.Type().(*types.PointerType)
			if !ok {
				c.exit("not a struct " + leftStruct.Type().String() + " " + expr.String())
			}
		}
		if ctypes.IsPointer(currentCandiceType) {
			currentCandiceType = currentCandiceType.(*ctypes.Pointer).Inner
		}
		if anonymous, ok := currentCandiceType.(*ctypes.Anonymous); ok {
			if anonymous.Modules != nil && len(anonymous.Modules) != 0 {
				module := anonymous.Modules[0]
				candiceType = c.modules[module].types[s.ElemType.Name()].candiceType.(*ctypes.Struct)
			} else {
				candiceType = c.types[s.ElemType.Name()].candiceType.(*ctypes.Struct)
			}
		} else {
			candiceType = currentCandiceType.(*ctypes.Struct)
		}
	}

	for {
		rightName, last := getName(expr.Right)
		i, field := candiceType.GetField(ast.RetrieveID(rightName))
		var inner types.Type
		inner = leftStruct.Type().(*types.PointerType).ElemType
		ptr := c.block().NewGetElementPtr(inner, leftStruct, zero, constant.NewInt(types.I32, int64(i)))
		leftStruct = ptr

		// Previously here we were unwrapping pointer struct values now we don't...
		// We only need to unwrap when we really need to load

		if last {
			break
		}

		expr = expr.Right.(*ast.BinaryOperation)
		candiceType = c.UnwrapStruct(field)
	}
	return leftStruct
}

func (c *Compiler) compileAdd(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
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
		return c.block().NewGetElementPtr(variable.Type().(*types.PointerType).ElemType, variable, zero, zero)
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
		//storage := c.block().NewAlloca(toReturnType)
		//c.block().NewStore(value, storage)
		return value
	}

	if ctypes.IsPointer(call.TypeParameters[0]) && types.IsPointer(variable.Type()) {
		c.doNotLoadIntoMemory = true
		return c.block().NewBitCast(variable, toReturnType)
	}
	c.exit("cant convert yet to this " + call.String() + "\n" + call.Parameters[0].GetType().String() + " " + variable.Type().String())
	panic("")
}
