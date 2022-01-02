package compiler

import (
	"bytes"
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"os/exec"
	"strings"
)

var zero value.Value = constant.NewInt(types.I32, 0)
var one value.Value = constant.NewInt(types.I32, 1)

type Compiler struct {
	errors                []error
	m                     *ir.Module
	blocks                []*ir.Block
	main                  *ir.Func
	types                 map[string]*Type
	definitions           map[string]value.Value
	builtins              map[string]func(*ast.BuiltinCall) value.Value
	definitionsToBePopped []string
	stacks                []map[string]value.Value
}

func New() *Compiler {
	m := ir.NewModule()
	main := m.NewFunc("main", types.I32)
	c := &Compiler{
		m:                     m,
		blocks:                []*ir.Block{main.NewBlock("_main")},
		definitions:           map[string]value.Value{},
		builtins:              map[string]func(*ast.BuiltinCall) value.Value{},
		types:                 map[string]*Type{},
		definitionsToBePopped: []string{"<>"},
		stacks:                []map[string]value.Value{map[string]value.Value{}},
	}
	c.initializeBuiltinLib()
	return c
}

/// Frequent private utils

// A small context stores variables stored by the current scope defined by if and for statements.
func (c *Compiler) createSmallContext() {
	c.definitionsToBePopped = append(c.definitionsToBePopped, "<>")
}

func (c *Compiler) addToCurrentSmallContext(variable string) {
	c.definitionsToBePopped = append(c.definitionsToBePopped, variable)
}

func (c *Compiler) removeCurrentSmallContext() {
	var i int
	stack := c.stack()
	for i = len(c.definitionsToBePopped) - 1; c.definitionsToBePopped[i] != "<>"; i-- {
		delete(stack, c.definitionsToBePopped[i])
	}
	c.definitionsToBePopped = c.definitionsToBePopped[:i]
}

func (c *Compiler) stack() map[string]value.Value {
	return c.stacks[len(c.stacks)-1]
}

func (c *Compiler) block() *ir.Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *Compiler) popBlock() *ir.Block {
	b := c.block()
	c.stacks = c.stacks[:len(c.stacks)-1]
	c.blocks = c.blocks[:len(c.blocks)-1]
	return b
}

func (c *Compiler) initializeBuiltinLib() {
	printf := c.m.NewFunc(
		"printf",
		types.I32,
		ir.NewParam("", types.NewPointer(types.I8)),
	)
	c.definitions["printf"] = printf
	printf.Sig.Variadic = true

	c.builtins["println"] = func(call *ast.BuiltinCall) value.Value {
		expressions := make([]value.Value, len(call.Parameters)+1)
		for i := 0; i < len(call.Parameters); i++ {
			expressions[i+1] = c.loadIfPointer(c.compileExpression(call.Parameters[i]))
		}
		constantString := strings.Builder{}
		// TODO: Here we would write a function that tries to do a toString()
		// 		for each expression
		for i := 0; i < len(call.Parameters); i++ {
			constantString.WriteString("%d ")
		}
		s := constantString.String()

		stringWithCharArrayType := constant.NewCharArrayFromString(s)

		// Define as global, we can keep it at all times on memory
		var globalDef value.Value

		if definition, ok := c.definitions[s]; !ok {
			globalDef = c.m.NewGlobalDef(s, stringWithCharArrayType)
			c.definitions[s] = globalDef
		} else {
			globalDef = definition
		}

		i8sType := c.block().NewGetElementPtr(
			// To be honest this is so strange, we are casting [i8 x len] to *[i8 x len]
			types.NewArray(uint64(len(s)), types.I8),
			globalDef,
			zero,
			zero,
		)
		expressions[0] = i8sType
		c.block().NewCall(c.definitions["printf"], expressions...)
		return constant.NewUndef(types.Void)
	}

	malloc := c.m.NewFunc(
		"malloc",
		types.NewPointer(types.I8),
		ir.NewParam("", types.I64),
	)
	c.definitions["malloc"] = malloc

	// alloc accepts one type parameter, and how many you want to allocate
	c.builtins["alloc"] = func(call *ast.BuiltinCall) value.Value {
		typeParameter := call.TypeParameters[0]
		toReturnType := types.NewPointer(c.ToLLVMType(typeParameter))
		length := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
		length = c.handleIntegerCast(types.I64, length)
		totalSize := c.block().NewMul(length, constant.NewInt(types.I64, typeParameter.SizeOf()))
		returnedValue := c.block().NewCall(c.definitions["malloc"], totalSize)
		castedValue := c.block().NewBitCast(returnedValue, toReturnType)
		alloca := c.block().NewAlloca(castedValue.Type())
		c.block().NewStore(castedValue, alloca)
		return alloca
	}

	c.builtins["cast"] = func(call *ast.BuiltinCall) value.Value {
		return c.handleCast(call)
	}

}

/// Public methods for the compiler

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

	switch t := tree.(type) {

	case *ast.ExpressionStatement:
		{
			c.compileExpression(t.Expression)
		}

	case *ast.StructStatement:
		{
			c.compileStruct(t)
		}

	case *ast.DeclarationStatement:
		{
			c.compileDeclaration(t)
		}

	case *ast.AssignmentStatement:
		{
			c.compileAssignment(t)
		}

	case *ast.Program:
		{
			for _, statement := range t.Statements {
				c.Compile(statement)
			}
			if len(c.blocks) == 1 {
				c.block().NewRet(constant.NewInt(types.I32, 0))
			}
			_ = c.popBlock()
		}

	}
}

func (c *Compiler) compileDeclaration(decl *ast.DeclarationStatement) {
	t := c.ToLLVMType(decl.Type)
	//if _, exists := c.types[t.Name()]; exists {
	//	if _, ok := c.types[t.Name()].candiceType.(*ctypes.Struct); ok {
	//		// Convert struct to pointer.
	//		// Reason why is that it's easy to copy a struct,
	//		// the main convention we should maintain is, a struct is always behind a pointer.
	//		// And when we copy make a manual copy before calling the function.
	//		t = types.NewPointer(t)
	//	}
	//}
	val := c.block().NewAlloca(t)
	c.block().NewStore(c.loadIfPointer(c.compileExpression(decl.Expression)), val)
	c.stack()[decl.Name] = val
	c.addToCurrentSmallContext(decl.Name)
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
	case *ast.Integer:
		{
			theType := c.ToLLVMType(e.Type)
			integerType := theType.(*types.IntType)
			return constant.NewInt(integerType, e.Value)
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
	}

	return nil
}

func (c *Compiler) compilePrefixExpression(prefix *ast.PrefixOperation) value.Value {
	prefixValue := c.compileExpression(prefix.Right)
	if prefix.Operation == ops.Subtract {
		prefixValue = c.loadIfPointer(prefixValue)
		return c.block().NewMul(prefixValue, constant.NewInt(prefixValue.Type().(*types.IntType), -1))
	}

	if prefix.Operation == ops.BinaryAND {
		allocatedValue := c.block().NewAlloca(prefixValue.Type())
		c.block().NewStore(prefixValue, allocatedValue)
		return allocatedValue
	}

	if prefix.Operation == ops.Multiply {
		prefixValue = c.block().NewLoad(prefixValue.Type().(*types.PointerType).ElemType, prefixValue)
		return prefixValue
	}

	return nil
}

// NOTE: change of plans, we are now loading identifiers stack references and if the caller needs it we
// load it there
func (c *Compiler) compileIdentifier(id *ast.Identifier) value.Value {
	return c.compileIdentifierReference(id)
}

func (c *Compiler) compileIdentifierReference(id *ast.Identifier) value.Value {
	identifier := c.stack()[id.Name]
	return identifier
}

/// Function calls
func (c *Compiler) compileBuiltinFunctionCall(ast *ast.BuiltinCall) value.Value {
	if fun, ok := c.builtins[ast.Name]; ok {
		return fun(ast)
	}
	panic("undefined builtin function @" + ast.Name)
}

func (c *Compiler) compileFunctionCall(ast *ast.Call) value.Value {
	return nil
}

func (c *Compiler) compileAssignment(assignment *ast.AssignmentStatement) {
	l := c.compileExpression(assignment.Left)
	r := c.loadIfPointer(c.compileExpression(assignment.Expression))
	c.block().NewStore(r, l)
}

func (c *Compiler) compileIndexAccess(access *ast.IndexAccess) value.Value {
	leftArray := c.loadIfPointer(c.compileExpression(access.Left))
	index := c.loadIfPointer(c.compileExpression(access.Access))
	pointer := c.block().NewGetElementPtr(leftArray.Type().(*types.PointerType).ElemType, leftArray, index)
	return pointer
}

func (c *Compiler) loadIfPointer(val value.Value) value.Value {
	if types.IsPointer(val.Type()) {
		return c.block().NewLoad(val.Type().(*types.PointerType).ElemType, val)
	}
	return val
}

func (c *Compiler) compileStructLiteral(strukt *ast.StructLiteral) value.Value {
	possibleStruct := c.types[strukt.Name]
	struktType, ok := possibleStruct.candiceType.(*ctypes.Struct)

	if !ok {
		panic(fmt.Sprintf("expected struct but got a %s", possibleStruct.candiceType))
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
/// Making redundant and easy to understand functions is better
/// than storing callbacks on a hashmap. Let's keep it simple.
func (c *Compiler) compileBinaryExpression(expr *ast.BinaryOperation) value.Value {
	switch expr.Operation {
	case ops.Multiply:
		{
			return c.compileMultiply(expr)
		}

	case ops.LessThanEqual:
		{

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
	}

	panic("unimplemented: " + expr.Operation.String())
	return nil
}

func getName(expr ast.Expression) (string, bool) {
	if bin, ok := expr.(*ast.BinaryOperation); ok {
		return bin.Left.(*ast.Identifier).Name, false
	}
	if bin, ok := expr.(*ast.Identifier); ok {
		return bin.Name, true
	}
	panic("?? " + expr.String())
}

func (c *Compiler) compileStructAccess(expr *ast.BinaryOperation) value.Value {
	leftStruct := c.compileExpression(expr.Left)
	var candiceType *ctypes.Struct
	if s, ok := leftStruct.Type().(*types.PointerType); ok {
		if types.IsPointer(s.ElemType) {
			leftStruct = c.loadIfPointer(leftStruct)
			s, ok = leftStruct.Type().(*types.PointerType)
			if !ok {
				panic("not a struct " + leftStruct.Type().String() + " " + expr.String())
			}
		}
		t := c.types[s.ElemType.Name()]
		candiceType, ok = t.candiceType.(*ctypes.Struct)
		if !ok {
			panic("not candice type ctypes struct: " + t.candiceType.String())
		}
	} else {
		panic("not a struct " + leftStruct.Type().String() + " " + expr.String())
	}
	for {
		rightName, last := getName(expr.Right)
		i, field := candiceType.GetField(rightName)
		var inner types.Type

		inner = leftStruct.Type().(*types.PointerType).ElemType

		ptr := c.block().NewGetElementPtr(inner, leftStruct, zero, constant.NewInt(types.NewInt(32), int64(i)))
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
	return c.block().NewAdd(leftValue, rightValue)
}

func (c *Compiler) compileMultiply(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewMul(leftValue, rightValue)
}

func (c *Compiler) compileSubtract(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	return c.block().NewSub(leftValue, rightValue)
}

func (c *Compiler) compileDivide(expr *ast.BinaryOperation) value.Value {
	leftValue := c.loadIfPointer(c.compileExpression(expr.Left))
	rightValue := c.loadIfPointer(c.compileExpression(expr.Right))
	if types.IsInt(leftValue.Type()) {
		return c.block().NewSDiv(leftValue, rightValue)
	}
	if types.IsFloat(leftValue.Type()) {
		panic("float arithmetic not implemented")
	}
	return nil
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
	variable := c.loadIfPointer(c.compileExpression(call.Parameters[0]))
	if ctypes.IsNumeric(call.TypeParameters[0]) && types.IsInt(variable.Type()) {
		return c.handleIntegerCast(toReturnType.(*types.IntType), variable)
	}
	panic("cant convert yet to this")
	return nil
}

func (c *Compiler) handleIntegerCast(toReturnType *types.IntType, variable value.Value) value.Value {
	variableType := variable.Type().(*types.IntType)
	if variableType.BitSize > toReturnType.BitSize {
		return c.block().NewTrunc(variable, toReturnType)
	}
	if variableType.BitSize == toReturnType.BitSize {
		return variable
	}
	return c.block().NewSExt(variable, toReturnType)
}
