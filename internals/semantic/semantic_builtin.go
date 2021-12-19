package semantic

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
)

func (s *Semantic) analyzeCast(castCall *ast.BuiltinCall) ctypes.Type {
	currentType := s.analyzeExpression(castCall.Parameters[0])
	toType := castCall.TypeParameters[0]
	if (ctypes.IsPointer(currentType) || ctypes.IsArray(currentType) || ctypes.IsNumeric(currentType)) &&
		(ctypes.IsPointer(toType) || ctypes.IsArray(toType) || ctypes.IsNumeric(toType)) {
		return toType
	}
	s.error("can't cast "+currentType.String()+" to "+toType.String(), castCall.Token)
	return ctypes.TODO()
}
