package compiler

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) createString(s string) value.Value {
	var globalDef value.Value
	s += string(byte(0))
	if definition, ok := c.globalBuiltinDefinitions[s]; !ok {
		globalDef = c.m.NewGlobalDef(s[:len(s)-1], constant.NewCharArrayFromString(s))
		c.globalBuiltinDefinitions[s] = globalDef
	} else {
		globalDef = definition
	}

	i8sType := c.block().NewGetElementPtr(
		// we are casting [i8 x len] to *i8
		types.NewArray(uint64(len(s)), types.I8),
		globalDef,
		zero,
		zero,
	)

	return i8sType
}
