package compiler

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) getFormatString(expressions []value.Value, t types.Type, call *ast.BuiltinCall, current int) string {
	if types.IsInt(t) {
		if integer, isUnsigned := call.Parameters[current].GetType().(*ctypes.UInteger); isUnsigned {
			if integer.BitSize > 32 {
				return "%llu "
			} else {
				return "%u "
			}
		} else if integer, isSigned := call.Parameters[current].GetType().(*ctypes.Integer); isSigned {
			if integer.BitSize > 32 {
				return "%lld "
			} else {
				return "%d "
			}
		}
	} else if pointer, isPointer := t.(*types.PointerType); isPointer {
		if _, ok := pointer.ElemType.(*types.IntType); ok {
			return "%s "
		} else {
			//
		}
	} else if float, isFloat := t.(*types.FloatType); isFloat {
		if float.Kind != types.FloatKindDouble {
			expressions[current+1] = c.handleFloatCast(types.Double, expressions[current+1])
		}
		return "%.3f "
	} else if strukt, isStruct := t.(*types.StructType); isStruct {
		expressions[current+1] = c.createString(ast.RetrieveID(strukt.TypeName))
		return "%s "
	} else if _, isPtr := t.(*types.PointerType); isPtr {
		return "*%p "
	}

	return ""
}
