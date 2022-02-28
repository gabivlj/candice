package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
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

func (c *Compiler) memcpy() *ir.Func {
	if fn, ok := c.globalBuiltinDefinitions["memcpy"]; ok {
		return fn.(*ir.Func)
	}
	memcpy := c.m.NewFunc("memcpy", types.Void, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I64))
	c.globalBuiltinDefinitions["memcpy"] = memcpy
	memcpy.CallingConv = enum.CallingConvC
	return memcpy
}

func (c *Compiler) strlen() *ir.Func {
	if fn, ok := c.globalBuiltinDefinitions["strlen"]; ok {
		return fn.(*ir.Func)
	}
	strlen := c.m.NewFunc("strlen", types.I64, ir.NewParam("", types.I8Ptr))
	c.globalBuiltinDefinitions["strlen"] = strlen
	strlen.CallingConv = enum.CallingConvC
	return strlen
}

func (c *Compiler) malloc() *ir.Func {
	if fn, ok := c.globalBuiltinDefinitions["malloc"]; ok {
		return fn.(*ir.Func)
	}
	malloc := c.m.NewFunc(
		"malloc",
		types.NewPointer(types.I8),
		ir.NewParam("", types.I64),
	)
	c.globalBuiltinDefinitions["malloc"] = malloc
	malloc.CallingConv = enum.CallingConvC
	return malloc
}

func (c *Compiler) strcmp() *ir.Func {
	if fn, ok := c.globalBuiltinDefinitions["strcmp"]; ok {
		return fn.(*ir.Func)
	}
	strcmp := c.m.NewFunc(
		"strcmp",
		types.I32,
		ir.NewParam("", types.I8Ptr),
		ir.NewParam("", types.I8Ptr),
	)
	c.globalBuiltinDefinitions["strcmp"] = strcmp
	strcmp.CallingConv = enum.CallingConvC
	return strcmp
}

// use this for concatenating strings
func (c *Compiler) concatenateMemoryI8(left value.Value, right value.Value) value.Value {
	strlen := c.strlen()
	len1 := c.block().NewCall(strlen, left)
	len2 := c.block().NewCall(strlen, right)
	totalLen := c.addOne(c.block().NewAdd(len1, len2))
	result := c.block().NewCall(c.malloc(), totalLen)
	memcpy := c.memcpy()
	c.block().NewCall(memcpy, result, left, len1)
	resultInt := c.block().NewIntToPtr(c.block().NewAdd(c.block().NewPtrToInt(result, types.I64), len1), types.I8Ptr)
	c.block().NewCall(memcpy, resultInt, right, c.addOne(len2))
	return result
}

// use this for comparing strings
func (c *Compiler) compareMemoryI8(left value.Value, right value.Value, pred enum.IPred) value.Value {
	return c.block().NewICmp(pred, c.block().NewCall(c.strcmp(), left, right), zero)
}
