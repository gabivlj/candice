package compiler

import (
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// toBool tries to pass this value to a i1 variable. If it's already a boolean it doesn't emit any bytecode.
func (c *Compiler) toBool(value value.Value) value.Value {
	if integer, ok := value.Type().(*types.IntType); ok {
		if integer.BitSize == 1 {
			return value
		}
		return c.block().NewICmp(enum.IPredNE, value, zero)
	}
	panic("can't pass to a boolean the value: " + value.String())
}

// handleIntegerCast tries to pass variable type to toReturnType
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
