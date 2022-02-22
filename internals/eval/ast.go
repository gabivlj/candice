package eval

import (
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
)

// SimplifyExpression tries to simplify an ast expression as much as possible,
// it tries to, for example, removing simple numeric casts.
func SimplifyExpression(expression ast.Expression) ast.Expression {
	switch t := expression.(type) {
	case *ast.BuiltinCall:
		{
			return simplifyBuiltinCall(t)
		}
	}

	return expression
}

func simplifyBuiltinCall(builtinCall *ast.BuiltinCall) ast.Expression {
	if builtinCall.Name == "cast" {
		return simplifyCastCall(builtinCall)
	}

	return builtinCall
}

func simplifyCastCall(castCall *ast.BuiltinCall) ast.Expression {
	t := castCall.TypeParameters[0]
	castCall.Parameters[0] = SimplifyExpression(castCall.Parameters[0])

	// Simplify case where we define variable where we make a cast:
	// integer := 1 as i64
	// >> integer := 1 (underlying type is i64)
	if ctypes.IsNumeric(t) && ctypes.IsNumeric(castCall.Parameters[0].GetType()) {
		integer, isParameterNumber := castCall.Parameters[0].(*ast.Integer)
		if isParameterNumber && !ctypes.IsFloat(t) {
			integer.Type = t
			return integer
		}
	}

	return castCall
}
