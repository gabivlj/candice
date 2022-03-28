package compiler

import (
	"fmt"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) handleComparisonOperations(expr *ast.BinaryOperation) value.Value {
	if left, isPointer := expr.Left.GetType().(*ctypes.Pointer); isPointer {
		if left.Inner == ctypes.I8 {
			return c.compareMemoryI8(c.loadIfPointer(c.compileExpression(expr.Left)),
				c.loadIfPointer(c.compileExpression(expr.Right)),
				c.getIPredComparison(expr.Operation, ctypes.I32),
			)
		}
	}

	if _, isFloat := expr.Left.GetType().(*ctypes.Float); isFloat {
		return c.block().NewFCmp(
			c.getFPredComparison(expr.Operation, expr.Left.GetType()),
			c.loadIfPointer(c.compileExpression(expr.Left)),
			c.loadIfPointer(c.compileExpression(expr.Right)),
		)
	}

	return c.block().NewICmp(c.getIPredComparison(expr.Operation, expr.Left.GetType()),
		c.loadIfPointer(c.compileExpression(expr.Left)),
		c.loadIfPointer(c.compileExpression(expr.Right)),
	)
}

func (c *Compiler) getFPredComparison(op ops.Operation, t ctypes.Type) enum.FPred {
	if _, ok := t.(*ctypes.Float); ok {
		switch op {
		case ops.GreaterThanEqual:
			return enum.FPredOGE
		case ops.GreaterThan:
			return enum.FPredOGT
		case ops.LessThan:
			return enum.FPredOLT
		case ops.LessThanEqual:
			return enum.FPredOLE
		case ops.Equals:
			return enum.FPredOEQ
		case ops.NotEquals:
			return enum.FPredONE
		}
	}

	c.exit("cant handle this type of float " + t.String() + " op: " + op.String() + fmt.Sprintf("%d", op))
	panic("")
}

func (c *Compiler) getIPredComparison(op ops.Operation, t ctypes.Type) enum.IPred {
	if _, ok := t.(*ctypes.Integer); ok {
		switch op {
		case ops.GreaterThanEqual:
			return enum.IPredSGE
		case ops.GreaterThan:
			return enum.IPredSGT
		case ops.LessThan:
			return enum.IPredSLT
		case ops.LessThanEqual:
			return enum.IPredSLE
		case ops.Equals:
			return enum.IPredEQ
		case ops.NotEquals:
			return enum.IPredNE
		}
	}

	if _, ok := t.(*ctypes.UInteger); ok {
		switch op {
		case ops.GreaterThanEqual:
			return enum.IPredUGE
		case ops.GreaterThan:
			return enum.IPredUGT
		case ops.LessThan:
			return enum.IPredULT
		case ops.LessThanEqual:
			return enum.IPredULE
		case ops.Equals:
			return enum.IPredEQ
		case ops.NotEquals:
			return enum.IPredNE
		}
	}

	if _, ok := t.(*ctypes.Pointer); ok {
		switch op {
		case ops.GreaterThanEqual:
			return enum.IPredSGE
		case ops.GreaterThan:
			return enum.IPredSGT
		case ops.LessThan:
			return enum.IPredSLT
		case ops.LessThanEqual:
			return enum.IPredSLE
		case ops.Equals:
			return enum.IPredEQ
		case ops.NotEquals:
			return enum.IPredNE
		}
	}

	c.exit("cant handle this type of integer " + t.String() + " op: " + op.String() + fmt.Sprintf("%d", op))
	panic("")
}
