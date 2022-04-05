package compiler

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"

	"github.com/gabivlj/candice/internals/ops"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func (c *Compiler) something() {
	// constant.
}

func (c *Compiler) handleConstantCast(val constant.Constant, originalType ctypes.Type, to types.Type, expr ast.Expression) constant.Constant {
	if types.IsInt(to) {
		to := to.(*types.IntType)
		if integerType, ok := val.Type().(*types.IntType); ok {
			if integerType.BitSize > to.BitSize {
				return constant.NewTrunc(val, to)
			}

			if integerType.BitSize < to.BitSize {
				if val.Type() == types.I1 {
					// SExt extends the bit sign of the val, i1 can be considered as a sign, so
					// it will be extended accross all the bits, for example if you cast
					// i1 '1' to i8, you will get 111...1110, which is -1.
					return constant.NewZExt(val, to)
				}

				return constant.NewSExt(val, to)
			}

			return val
		}

		if _, ok := val.Type().(*types.FloatType); ok {
			if ctypes.IsUnsignedInteger(originalType) {
				return constant.NewFPToUI(val, to)
			}

			return constant.NewFPToSI(val, to)
		}

		c.exitErrorExpression("can't handle this cast", expr)
	}

	if types.IsFloat(to) {
		to := to.(*types.FloatType)
		if _, ok := val.Type().(*types.IntType); ok {
			if !ctypes.IsUnsignedInteger(originalType) {
				return constant.NewSIToFP(val, to)
			}

			return constant.NewUIToFP(val, to)
		}

		if valFloat, ok := val.Type().(*types.FloatType); ok {
			if valFloat.Kind == types.FloatKindDouble && to.Kind == types.FloatKindFloat {
				return constant.NewFPTrunc(val, to)
			}

			if valFloat.Kind == to.Kind {
				return val
			}

			return constant.NewFPExt(val, to)
		}
	}

	c.exitErrorExpression("can't handle this cast", expr)
	return nil
}

func (c *Compiler) compileConstantExpression(expr ast.Expression) constant.Constant {
	switch expr := expr.(type) {
	case *ast.Identifier:
		return c.compileIdentifier(expr).(constant.Constant)
	case *ast.BinaryOperation:
		left := c.compileConstantExpression(expr.Left)
		right := c.compileConstantExpression(expr.Left)
		return c.compileOperation(expr.Operation, expr, left, right)
	case *ast.Integer:
		return c.compileExpression(expr).(constant.Constant)
	case *ast.Float:
		return c.compileExpression(expr).(constant.Constant)
	case *ast.StringLiteral:
		return constant.NewCharArrayFromString(expr.Value + string(byte(0)))
	case *ast.BuiltinCall:
		if expr.Name == "cast" {
			return c.handleConstantCast(
				c.compileConstantExpression(expr.Parameters[0]),
				expr.Parameters[0].GetType(),
				c.ToLLVMType(expr.GetType()),
				expr,
			)
		}
		c.exitErrorExpression("can't handle this constant builtin call", expr)
	default:
		c.exitErrorExpression("unexpected non constant operation", expr)
		panic(nil)
	}
	c.exitErrorExpression("can't handle this constant expression", expr)
	panic(nil)
}

func (c *Compiler) compileOperation(op ops.Operation, expr *ast.BinaryOperation, first, second constant.Constant) constant.Constant {
	switch op {
	case ops.Add:
		if types.IsFloat(first.Type()) {
			return constant.NewFAdd(first, second)
		}

		return constant.NewAdd(first, second)

	case ops.Subtract:
		if types.IsFloat(first.Type()) {
			return constant.NewFSub(first, second)
		}

		return constant.NewSub(first, second)

	case ops.Multiply:
		if types.IsFloat(first.Type()) {
			return constant.NewFMul(first, second)
		}

		return constant.NewMul(first, second)

	case ops.BinaryAND:
		return constant.NewAnd(first, second)
	case ops.BinaryXOR:
		return constant.NewXor(first, second)
	case ops.BinaryOR:
		return constant.NewOr(first, second)
	case ops.LeftShift:
		return constant.NewShl(first, second)
	case ops.RightShift:
		return constant.NewLShr(first, second)
	case ops.Divide:
		if types.IsInt(first.Type()) {
			if _, isUnsigned := expr.Type.(*ctypes.UInteger); isUnsigned {
				return constant.NewUDiv(first, second)
			}
			return constant.NewSDiv(first, second)
		}
		if types.IsFloat(first.Type()) {
			return constant.NewFDiv(first, second)
		}
	case ops.Modulo:
		if types.IsInt(first.Type()) {
			if _, isUnsigned := expr.Type.(*ctypes.UInteger); isUnsigned {
				return constant.NewURem(first, second)
			}
			return constant.NewSRem(first, second)
		}
		if types.IsFloat(first.Type()) {
			return constant.NewFRem(first, second)
		}
	}

	c.exitInternalError("can't handle operation " + op.String() + " on " + first.String() + " and " + second.String())
	return nil
}
