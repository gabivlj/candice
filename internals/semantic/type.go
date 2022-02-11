package semantic

import "github.com/gabivlj/candice/internals/ctypes"

type SemanticType struct {
	parentFunction *ctypes.Function
	Type           ctypes.Type
}

func (s *Semantic) newType(t ctypes.Type) *SemanticType {
	return &SemanticType{
		Type:           t,
		parentFunction: s.currentFunctionBeingAnalyzed,
	}
}
