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
		castCall.Type = toType
		return toType
	}
	s.error("can't cast "+currentType.String()+" to "+toType.String(), castCall.Token)
	return ctypes.TODO()
}

func (s *Semantic) analyzeAlloc(allocCall *ast.BuiltinCall) ctypes.Type {
	t := allocCall.TypeParameters[0]
	s.replaceAnonymous(t)
	expr := s.analyzeExpression(allocCall.Parameters[0])
	if !ctypes.IsNumeric(expr) {
		s.typeMismatchError(allocCall.String(), allocCall.Token, ctypes.I32, expr)
	}
	allocCall.Type = &ctypes.Pointer{Inner: t}
	return allocCall.Type
}

func (s *Semantic) analyzePrintln(_ *ast.BuiltinCall) ctypes.Type {
	return ctypes.VoidType
}

func (s *Semantic) analyzeFree(freeCall *ast.BuiltinCall) ctypes.Type {
	if len(freeCall.Parameters) != 1 {
		s.error("expected one parameter for free builtin call", freeCall.Token)
		return ctypes.TODO()
	}
	if !ctypes.IsPointer(s.analyzeExpression(freeCall.Parameters[0])) {
		s.error("expected pointer type for free call", freeCall.Token)
	}
	return ctypes.VoidType
}
