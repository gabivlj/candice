package compiler

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"os/exec"
)

type Compiler struct {
	errors []error
	m *ir.Module
	blocks []*ir.Block
	main *ir.Func
}


func New() *Compiler {
	m := ir.NewModule()
	main := m.NewFunc("main", types.I32)
	return &Compiler{
		m: m,
		blocks: []*ir.Block{main.NewBlock("_main")},
	}
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

/// Public methods for the compiler

// Execute generates and executes the executable
func (c *Compiler) Execute() error {
	err := GenerateExecutable(c.m, "exec")
	if err != nil {
		return err
	}
	err = exec.Command("./exec").Run()
	return err
}

// Compile compiles the entire ast
// It makes weak type checks, it will assume that the returned types
// are right. Usually you would want to semantically check the tree before
// calling this.
func (c *Compiler) Compile(tree ast.Node) {

	switch t := tree.(type) {

	case *ast.ExpressionStatement: {
		c.compileExpression(t.Expression)
	}

	case *ast.Program: {
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
	case *ast.Integer: {
		return constant.NewInt(types.I64, e.Value)
	}

	case *ast.BinaryOperation: {
		return c.compileBinaryExpression(e)
	}
	}

	return nil
}

/// Simple binary compilations
/// Making redundant and easy to understand functions is better
/// than storing callbacks on a hashmap. Let's keep it simple.

func (c *Compiler) compileBinaryExpression(expr *ast.BinaryOperation) value.Value {
	switch expr.Operation {
	case ops.Multiply: {
		return c.compileMultiply(expr)
	}

	case ops.LessThanEqual: {

	}

	case ops.Plus: {
		return c.compileAdd(expr)
	}

	case ops.Divide: {
		return c.compileDivide(expr)
	}

	case ops.Minus: {
		return c.compileSubtract(expr)
	}

	case ops.BinaryXOR: {
		return c.compileXorBinary(expr)
	}

	case ops.BinaryAND: {
		return c.compileAndBinary(expr)
	}

	case ops.BinaryOR: {
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
