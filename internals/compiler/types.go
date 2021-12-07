package compiler

import (
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/llir/llvm/ir/types"
)

type Type struct {
	llvmType    types.Type
	candiceType ctypes.Type
}

func (c *Compiler) ToLLVMType(t ctypes.Type) types.Type {
	switch el := t.(type) {
	case *ctypes.Integer:
		{
			return types.NewInt(uint64(el.BitSize))
		}
	case *ctypes.UInteger:
		{
			return types.NewInt(uint64(el.BitSize))
		}
	case *ctypes.Pointer:
		{
			return types.NewPointer(c.ToLLVMType(el.Inner))
		}
	case *ctypes.Struct:
		{

			llvmTypes := make([]types.Type, len(el.Fields))
			for i, field := range el.Fields {
				llvmTypes[i] = c.ToLLVMType(field)
			}
			s := types.NewStruct(llvmTypes...)
			return s
		}
	case *ctypes.Array:
		{
			return types.NewArray(uint64(el.Length), c.ToLLVMType(el.Inner))
		}
	case *ctypes.Void:
		{
			return types.Void
		}
	case *ctypes.Float:
		{
			return types.Float
		}
	case *ctypes.Anonymous:
		{
			return c.types[el.Name].llvmType
		}
	}
	return nil
}
