package compiler

import (
	"github.com/gabivlj/candice/internals/ctypes"
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
	if _, ok := value.Type().(*types.PointerType); ok {
		return c.block().NewICmp(enum.IPredNE, c.block().NewPtrToInt(value, types.I64), zero)
	}
	c.exitInternalError("can't pass to a boolean the value: " + value.String())
	panic("")
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

// performs a cast bitsize, if first parameter is a float the second param should be a float as well
// if you call this function and toReturnType is not
// a float or an integer the program will probably panic
func (c *Compiler) handleNumericBitCast(toReturnType types.Type, variable value.Value) value.Value {
	integerType, isInteger := toReturnType.(*types.IntType)
	if isInteger {
		return c.handleIntegerCast(integerType, variable)
	}
	return c.handleFloatCast(toReturnType.(*types.FloatType), variable)
}

func (c *Compiler) handleFloatCast(toReturnType *types.FloatType, variable value.Value) value.Value {
	variableType := variable.Type().(*types.FloatType)
	if variableType.Kind == types.FloatKindDouble && toReturnType.Kind == types.FloatKindFloat {
		return c.block().NewFPTrunc(variable, toReturnType)
	}
	if variableType.Kind == toReturnType.Kind {
		return variable
	}

	return c.block().NewFPExt(variable, toReturnType)
}

func (c *Compiler) handleFloatIntCast(typeParameter ctypes.Type, parameter ctypes.Type, variable value.Value, toReturn types.Type) value.Value {

	if ctypes.IsFloat(parameter) {
		return c.floatToInt(toReturn.(*types.IntType), variable, ctypes.IsUnsignedInteger(typeParameter))
	} else if ctypes.IsFloat(typeParameter) {
		return c.intToFloat(toReturn.(*types.FloatType), variable, ctypes.IsUnsignedInteger(parameter))
	}

	c.exitInternalError("atleast one parameter need to be a float")
	panic("")
}

func (c *Compiler) floatToInt(toReturn *types.IntType, variable value.Value, isUnsigned bool) value.Value {
	if isUnsigned {
		return c.block().NewFPToUI(variable, toReturn)
	}
	return c.block().NewFPToSI(variable, toReturn)
}

func (c *Compiler) intToFloat(toReturn *types.FloatType, variable value.Value, isUnsigned bool) value.Value {
	if isUnsigned {
		return c.block().NewUIToFP(variable, toReturn)
	}

	return c.block().NewSIToFP(variable, toReturn)
}
