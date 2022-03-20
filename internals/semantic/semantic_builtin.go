package semantic

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
)

func (s *Semantic) analyzeCast(castCall *ast.BuiltinCall) ctypes.Type {
	currentType := s.UnwrapAnonymous(s.analyzeExpression(castCall.Parameters[0]))
	toType := s.UnwrapAnonymous(castCall.TypeParameters[0])
	castCall.TypeParameters[0] = toType
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
		s.typeMismatchError(allocCall.String(), allocCall.Parameters[0], allocCall.Token, ctypes.I32, expr)
	}
	allocCall.Type = &ctypes.Pointer{Inner: t}
	return allocCall.Type
}

func (s *Semantic) analyzeUnreachable(_ *ast.BuiltinCall) ctypes.Type {
	s.returns = true
	return s.currentExpectedReturnType
}

func (s *Semantic) analyzePrintln(printCall *ast.BuiltinCall) ctypes.Type {
	for _, param := range printCall.Parameters {
		s.analyzeExpression(param)
	}
	return ctypes.VoidType
}

func (s *Semantic) analyzeFree(freeCall *ast.BuiltinCall) ctypes.Type {
	if len(freeCall.Parameters) != 1 {
		s.error("expected one parameter for free builtin call", freeCall.Token)
		return ctypes.TODO()
	}
	if !ctypes.IsPointer(s.UnwrapAnonymous(s.analyzeExpression(freeCall.Parameters[0]))) {
		s.error("expected pointer type for free call", freeCall.Token)
	}
	return ctypes.VoidType
}

func (s *Semantic) analyzeRealloc(reallocCall *ast.BuiltinCall) ctypes.Type {
	if len(reallocCall.Parameters) != 2 {
		s.error("expected two parameters for realloc builtin call", reallocCall.Token)
		return ctypes.TODO()
	}

	t := s.UnwrapAnonymous(s.analyzeExpression(reallocCall.Parameters[0]))

	if !ctypes.IsPointer(t) {
		s.error("expected pointer type for realloc call", reallocCall.Token)
	}

	secondParameter := s.analyzeExpression(reallocCall.Parameters[1])

	if _, isInteger := secondParameter.(*ctypes.Integer); !isInteger {
		s.typeMismatchError(reallocCall.String(), reallocCall.Parameters[1], reallocCall.Token, ctypes.I64, secondParameter)
	}

	reallocCall.Type = t

	return t
}

func (s *Semantic) analyzeSizeOf(sizeOfCall *ast.BuiltinCall) ctypes.Type {
	if len(sizeOfCall.TypeParameters) != 1 {
		s.error("expected one type parameter for sizeOf builtin call", sizeOfCall.Token)
		return ctypes.TODO()
	}
	sizeOfCall.TypeParameters[0] = s.UnwrapAnonymous(sizeOfCall.TypeParameters[0])
	sizeOfCall.Type = sizeOfCall.TypeParameters[0]
	return ctypes.I32
}

func (s *Semantic) analyzeAsm(asmCall *ast.BuiltinCall) ctypes.Type {
	if len(asmCall.Parameters) == 0 {
		s.error("expected a constant string literal on asm builtin call", asmCall.Token)
		return ctypes.TODO()
	}

	_, isString := asmCall.Parameters[0].(*ast.StringLiteral)
	if !isString {
		s.error("expected a constant string literal on asm builtin call, got: "+asmCall.Parameters[0].String(), asmCall.Token)
		return ctypes.TODO()
	}

	return s.UnwrapAnonymous(asmCall.TypeParameters[0])
}
