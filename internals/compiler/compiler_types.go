package compiler

import (
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Type struct {
	llvmType    types.Type
	candiceType ctypes.Type
}

type Value struct {
	Value    value.Value
	Type     ctypes.Type
	Constant bool
}

func (c *Compiler) searchForType(name string) types.Type {
	if t, ok := c.types[name]; ok {
		return t.llvmType
	}

	// TODO: Optimize this to use a single lookup
	// we need to do this because we don't know the module name from a struct
	// and sometimes we will find imported types from other modules that have been parsed.
	for _, module := range c.modules {
		if t, ok := module.types[name]; ok {
			return t.llvmType
		}
	}

	return nil
}

// UnwrapFieldAccessor unwraps the underlying pointer/anonymous/struct/union type and returns a pure field accessor.
// This is useful when you have a pointer to a struct/union, but you want to get some fields.
func (c *Compiler) UnwrapFieldAccessor(field ctypes.Type) ctypes.FieldType {
	prev, ok := field.(*ctypes.Pointer)
	var candiceType ctypes.FieldType
	if ok {
		possibleStruct, ok := prev.Inner.(ctypes.FieldType)
		if !ok {
			candiceType = c.types[prev.Inner.(*ctypes.Anonymous).Name].candiceType.(ctypes.FieldType)
		} else {
			candiceType = possibleStruct
		}
	} else {
		possibleStruct, ok := field.(ctypes.FieldType)
		if !ok {
			candiceType = c.types[field.(*ctypes.Anonymous).Name].candiceType.(ctypes.FieldType)
		} else {
			candiceType = possibleStruct
		}

	}
	return candiceType
}

func (c *Compiler) retrieveInnerAnonymousAndUnwrap(field ctypes.Type) ctypes.Type {
	field, _ = ctypes.UnwrapPossiblePointerAndDepth(field)
	return c.context.UnwrapAnonymous(field)
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

	case *ctypes.Union:
		{
			if union := c.searchForType(el.Name); union != nil {
				return union
			}

			unionType := types.NewArray(uint64(el.SizeOf()), types.I8)
			return unionType
		}

	case *ctypes.Struct:
		{
			if strukt := c.searchForType(el.Name); strukt != nil {
				return strukt
			}

			llvmTypes := make([]types.Type, 0, len(el.Fields))
			s := types.NewStruct()
			s.Opaque = true
			c.types[el.Name] = &Type{llvmType: s, candiceType: ctypes.TODO()}
			for _, field := range el.Fields {
				llvmTypes = append(llvmTypes, c.ToLLVMType(field))
			}
			s.Fields = llvmTypes
			s.Opaque = false
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
			if el.BitSize >= 64 {
				return types.Double
			}

			if el.BitSize >= 128 {
				return types.FP128
			}

			return types.Float
		}
	case *ctypes.Anonymous:
		{
			t := c.types[el.Name]

			// If type doesn't exist and the anonymous type references
			// other modules, let's find it on already compiled modules.
			if t == nil && el.Modules != nil && len(el.Modules) > 0 {
				l := c.modules[el.Modules[0]].ToLLVMType(el)
				return l
			}
			return t.llvmType
		}

	case *ctypes.Function:
		{
			returnType := c.ToLLVMType(el.Return)
			parameters := make([]types.Type, 0, len(el.Parameters))
			for _, param := range el.Parameters {
				parameters = append(parameters, c.ToLLVMType(param))
			}
			function := types.NewPointer(types.NewFunc(returnType, parameters...))
			return function
		}
	}

	c.exitInternalError("can't convert to LLVM type: " + t.String())
	panic("")
}
