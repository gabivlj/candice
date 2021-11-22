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

func (c *Compiler) Execute() error {
	err := GenerateExecutable(c.m, "exec")
	if err != nil {
		return err
	}
	err = exec.Command("./exec").Run()
	return err
}

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


func (c *Compiler) block() *ir.Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *Compiler) popBlock() *ir.Block {
	b := c.block()
	c.blocks = c.blocks[:len(c.blocks)-1]
	return b
}