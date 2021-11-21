package ast

import (
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/pkg/a"
	"testing"
)

func TestIdentifier_String(t *testing.T) {
	identifier := Identifier{Name: "Cool"}
	a.Assert(identifier.String() == "Cool")
}

func TestBinaryOperation_String(t *testing.T) {
	const integer = 3
	const integerRight = 4
	binaryOp := BinaryOperation{
		Operation: ops.Plus,
		Left: &BinaryOperation{
			Operation: ops.Multiply,
			Left: &Integer{Value: integer}, Right: &Integer{Value: integerRight},
		},
		Right: &Integer{Value: integer},
	}
	a.Assert(
		binaryOp.String() ==
			`((3 * 4) + 3)`,
	)
}

