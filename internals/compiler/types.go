package compiler

import (
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/llir/llvm/ir/types"
)

type Type struct {
	llvmType    types.Type
	candiceType ctypes.Type
}

// UnwrapStruct unwraps the underlying pointer/anonymous/struct type and returns a pure struct.
// This is useful when you have a pointer to a struct, but you want to get some fields.
func (c *Compiler) UnwrapStruct(field ctypes.Type) *ctypes.Struct {
	prev, ok := field.(*ctypes.Pointer)
	var candiceType *ctypes.Struct
	if ok {
		possibleStruct, ok := prev.Inner.(*ctypes.Struct)
		if !ok {
			candiceType = c.types[prev.Inner.(*ctypes.Anonymous).Name].candiceType.(*ctypes.Struct)
		} else {
			candiceType = possibleStruct
		}
	} else {
		possibleStruct, ok := field.(*ctypes.Struct)
		if !ok {
			candiceType = c.types[field.(*ctypes.Anonymous).Name].candiceType.(*ctypes.Struct)
		} else {
			candiceType = possibleStruct
		}

	}
	return candiceType
}

// GetPureStruct tries to check if the type is a struct or an anonymous type that in the
// type repository is a struct. Doesn't try to unwrap the underlying type.
func (c *Compiler) GetPureStruct(t ctypes.Type) *ctypes.Struct {
	if strukt, ok := t.(*ctypes.Struct); !ok {
		if an, ok := t.(*ctypes.Anonymous); ok {
			if strukt, ok := c.types[an.Name].candiceType.(*ctypes.Struct); ok {
				return strukt
			}
		}
		return nil
	} else {
		return strukt
	}
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