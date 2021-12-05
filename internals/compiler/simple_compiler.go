package compiler

import (
	"bytes"
	"github.com/gabivlj/candice/internals/ast"
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
	errors      []error
	m           *ir.Module
	blocks      []*ir.Block
	main        *ir.Func
	definitions map[string]value.Value
	builtins    map[string]func(*ast.BuiltinCall) value.Value
}

func New() *Compiler {
	m := ir.NewModule()
	main := m.NewFunc("main", types.I32)
	c := &Compiler{
		m:           m,
		blocks:      []*ir.Block{main.NewBlock("_main")},
		definitions: map[string]value.Value{},
		builtins: map[string]func(*ast.BuiltinCall) value.Value{},
	}
	c.initializeBuiltinLib()
	return c
}

/// Frequent private utils

func (c *Compiler) block() *ir.Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *Compiler) popBlock() *ir.Block {
	b := c.block()
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
			expressions[i+1] = c.compileExpression(call.Parameters[i])
		}
		constantString := strings.Builder{}
		// TODO: Here we would write a function that tries to do a toString()
		// 		for each expression
		for i := 0; i < len(call.Parameters); i++ {
			constantString.WriteString("%d ")
		}
		s := constantString.String()
		stringWithCharArrayType := constant.NewCharArrayFromString(s)
		i8sType := c.block().NewGetElementPtr(
			// To be honest this is so strange, we are casting [i8 x len] to *[i8 x len]
			types.NewArray(uint64(len(s)), types.I8),
			c.m.NewGlobalDef(s, stringWithCharArrayType),
			zero,
			zero,
		)
		expressions[0] = i8sType
		c.block().NewCall(c.definitions["printf"], expressions...)
		return constant.NewUndef(types.Void)
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

func (c *Compiler) compileExpression(expression ast.Expression) value.Value {

	switch e := expression.(type) {
	case *ast.Integer:
		{
			return constant.NewInt(types.I64, e.Value)
		}

	case *ast.BinaryOperation:
		{
			return c.compileBinaryExpression(e)
		}

	case *ast.Call:
		{
			return c.compileFunctionCall(e)
		}

	case *ast.BuiltinCall:
		{
			c.compileBuiltinFunctionCall(e)
		}
	}

	return nil
}

/// Identifier
func (c *Compiler) compileIdentifier(ast *ast.Identifier) value.Value {
	return nil
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

	case ops.Plus:
		{
			return c.compileAdd(expr)
		}

	case ops.Divide:
		{
			return c.compileDivide(expr)
		}

	case ops.Minus:
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

	}

	panic("unimplemented: " + expr.Operation.String())
	return nil
}

func (c *Compiler) compileAdd(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewAdd(leftValue, rightValue)
}

func (c *Compiler) compileMultiply(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewMul(leftValue, rightValue)
}

func (c *Compiler) compileSubtract(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewSub(leftValue, rightValue)
}

func (c *Compiler) compileDivide(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	if types.IsInt(leftValue.Type()) {
		return c.block().NewSDiv(leftValue, rightValue)
	}
	if types.IsFloat(leftValue.Type()) {
		panic("float arithmetic not implemented")
	}
	return nil
}

func (c *Compiler) compileAndBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewAnd(leftValue, rightValue)
}

func (c *Compiler) compileOrBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewOr(leftValue, rightValue)
}

func (c *Compiler) compileXorBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewXor(leftValue, rightValue)
}

func (c *Compiler) compileShiftRightBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewLShr(leftValue, rightValue)
}

func (c *Compiler) compileShiftLeftBinary(expr *ast.BinaryOperation) value.Value {
	leftValue := c.compileExpression(expr.Left)
	rightValue := c.compileExpression(expr.Right)
	return c.block().NewShl(leftValue, rightValue)
}
