package semantic

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/node"
	"github.com/gabivlj/candice/internals/ops"
)

type SemanticType struct {
	parentFunction *ctypes.Function
	Type           ctypes.Type
	IsConstant     bool
}

func (s *Semantic) newType(t ctypes.Type) *SemanticType {
	return &SemanticType{
		Type:           t,
		parentFunction: s.currentFunctionBeingAnalyzed,
	}
}

func (s *Semantic) wrapInOperators(expression ast.Expression, times int, operator ops.Operation) ast.Expression {
	newType := expression.GetType()
	if operator == ops.Multiply {
		ptr, ok := newType.(*ctypes.Pointer)
		if !ok {
			s.errorWithExpression("can't unwrap non pointer", expression)
			return expression
		}

		newType = ptr.Inner
	} else if operator == ops.Reference {
		newType = &ctypes.Pointer{Inner: newType}
	}

	for i := 0; i < times; i++ {
		expression = &ast.PrefixOperation{
			Operation: operator,
			Node: &node.Node{
				Type:  newType,
				Token: expression.GetToken(),
			},
			Right: expression,
		}
	}
	return expression
}
