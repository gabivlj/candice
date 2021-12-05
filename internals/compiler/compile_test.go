package compiler

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/pkg/a"
	"testing"
)

func TestCompiler_CompileExpression_With_AddSubtractMultiplyDivide(t *testing.T) {
	c := New()
	c.Compile(
		&ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Expression: &ast.BuiltinCall{
						Name: "println",
						Parameters: []ast.Expression{
							&ast.BinaryOperation{
								Operation: ops.Multiply,
								Left: &ast.BinaryOperation{
									Left: &ast.BinaryOperation{
										Operation: ops.Plus,
										Left:      &ast.Integer{Value: 3},
										Right:     &ast.Integer{Value: 3},
									},
									Right: &ast.BinaryOperation{
										Operation: ops.Divide,
										Left: &ast.BinaryOperation{
											Operation: ops.Minus,
											Left:      &ast.Integer{Value: 332},
											Right:     &ast.Integer{Value: 1},
										},
										Right: &ast.Integer{Value: 3},
									},
								},
								Right: &ast.Integer{Value: 5},
							},
						},
					},
				},
			},
		},
	)

	a.AssertEqual("580", string(a.UnwrapBytes(c.Execute())))
}

func TestCompiler_CompileExpression_With_And(t *testing.T) {
	c := New()
	c.Compile(
		&ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Expression: &ast.BuiltinCall{
						Name: "println",
						Parameters: []ast.Expression{
							&ast.BinaryOperation{
								Operation: ops.BinaryAND,
								Left: &ast.Integer{Value: 3},
								Right: &ast.Integer{Value: 5},
							},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3 & 5), string(a.UnwrapBytes(c.Execute())))
}

func TestCompiler_CompileExpression_With_Or(t *testing.T) {
	c := New()
	c.Compile(
		&ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Expression: &ast.BuiltinCall{
						Name: "println",
						Parameters: []ast.Expression{
							&ast.BinaryOperation{
								Operation: ops.BinaryOR,
								Left: &ast.Integer{Value: 3},
								Right: &ast.Integer{Value: 5},
							},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3 | 5), string(a.UnwrapBytes(c.Execute())))
}

func TestCompiler_CompileExpression_With_Xor(t *testing.T) {
	c := New()
	c.Compile(
		&ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Expression: &ast.BuiltinCall{
						Name: "println",
						Parameters: []ast.Expression{
							&ast.BinaryOperation{
								Operation: ops.BinaryXOR,
								Left: &ast.Integer{Value: 3322323},
								Right: &ast.Integer{Value: 51231212},
							},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3322323 ^ 51231212), string(a.UnwrapBytes(c.Execute())))
}
