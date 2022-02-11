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

	if _, isFloat := expr.Type.(*ctypes.Float); isFloat {
		panic("can't use floats yet")
	}

	return c.block().NewICmp(c.getIPredComparison(expr.Operation, expr.Left.GetType()),
		c.loadIfPointer(c.compileExpression(expr.Left)),
		c.loadIfPointer(c.compileExpression(expr.Right)))
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

	panic("cant handle this type of integer " + t.String() + " op: " + op.String() + fmt.Sprintf("%d", op))
}
