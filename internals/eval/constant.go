package eval

import (
	"github.com/gabivlj/candice/internals/ast"
)

func EvaluateConstantExpression(expression ast.Expression) Value {
	switch typedExpression := expression.(type) {
	case *ast.Integer:
		{
			return &Integer{typedExpression.Value}
		}

	case *ast.Identifier:
		{
			value, ok := constants[ast.RetrieveID(typedExpression.Name)]
			if !ok {
				return &Error{
					Message: "unknown identifier for constant expression" + expression.String(),
					Token:   expression.GetToken(),
				}
			}

			return value
		}
	}

	return &Error{
		Message: "can't evaluate expression " + expression.String(),
		Token:   expression.GetToken(),
	}
}
