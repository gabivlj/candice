package compiler

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
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
								Left:      &ast.Integer{Value: 3},
								Right:     &ast.Integer{Value: 5},
							},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3&5), string(a.UnwrapBytes(c.Execute())))
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
								Left:      &ast.Integer{Value: 3},
								Right:     &ast.Integer{Value: 5},
							},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3|5), string(a.UnwrapBytes(c.Execute())))
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
								Left:      &ast.Integer{Value: 3322323},
								Right:     &ast.Integer{Value: 51231212},
							},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3322323^51231212), string(a.UnwrapBytes(c.Execute())))
}

func TestCompiler_CompileStruct(t *testing.T) {
	/*
			This is an insane test to type as an AST.

			what we are testing is:

			struct Point2 {
				x int
				y int
			}

			struct Point {
				x int
				y int
				self Point2
			}

		    point := Point{ ..., self: { x = 3} }
			assertEqual(point.self.x, 3)
	*/
	c := New()
	pointStruct := &ast.StructStatement{
		Type: &ctypes.Struct{
			Fields: []ctypes.Type{
				&ctypes.Integer{BitSize: 64},
				&ctypes.Integer{BitSize: 64},
				&ctypes.Anonymous{Name: "Point2"},
			},
			Names: []string{"x", "y", "self"},
			Name:  "Point",
		},
	}
	point2Struct := &ast.StructStatement{
		Type: &ctypes.Struct{
			Fields: []ctypes.Type{
				&ctypes.Integer{BitSize: 64},
				&ctypes.Integer{BitSize: 64},
			},
			Names: []string{"x", "y"},
			Name:  "Point2",
		},
	}
	c.Compile(&ast.Program{
		Statements: []ast.Statement{
			point2Struct,
			pointStruct,
			&ast.DeclarationStatement{
				Name: "point",
				Type: &ctypes.Anonymous{Name: "Point"},
				Expression: &ast.StructLiteral{
					Name: "Point",
					Values: []ast.StructValue{
						{
							Name:       "x",
							Expression: &ast.Integer{Value: 3},
						},
						{
							Name:       "y",
							Expression: &ast.Integer{Value: 3},
						},
						{
							Name: "self",
							Expression: &ast.StructLiteral{
								Name: "Point2",
								Values: []ast.StructValue{
									{
										Name:       "x",
										Expression: &ast.Integer{Value: 3},
									},
									{
										Name:       "y",
										Expression: &ast.Integer{Value: 3},
									},
								},
							},
						},
					},
				},
			},
			&ast.DeclarationStatement{
				Name: "x",
				Type: &ctypes.Integer{BitSize: 64},
				Expression: &ast.BinaryOperation{
					Left: &ast.Identifier{Name: "point"},
					Right: &ast.BinaryOperation{
						Left:      &ast.Identifier{Name: "self"},
						Right:     &ast.Identifier{Name: "x"},
						Operation: ops.Dot,
					},
					Operation: ops.Dot,
				},
			},
			&ast.ExpressionStatement{Expression: &ast.BuiltinCall{
				Name: "println",
				Parameters: []ast.Expression{&ast.Identifier{
					Name: "x",
				}},
			}},
		},
	})
	a.AssertEqual(fmt.Sprintf("%d", 3), string(a.UnwrapBytes(c.Execute())))
}

func TestCompiler_CompileExpression_With_Malloc(t *testing.T) {
	c := New()
	c.Compile(&ast.Program{
		Statements: []ast.Statement{
			&ast.DeclarationStatement{
				Name: "coolStuff",
				Type: &ctypes.Pointer{Inner: &ctypes.Integer{BitSize: 32}},
				Expression: &ast.BuiltinCall{
					Name:           "alloc",
					TypeParameters: []ctypes.Type{&ctypes.Integer{BitSize: 32}},
					Parameters:     []ast.Expression{&ast.Integer{Value: 5}},
				},
			},
		},
	})
	_, err := c.Execute()
	a.AssertErr(err)
}

func TestCompiler_CompileExpression_With_Sum_And_Decl(t *testing.T) {
	c := New()
	binOp := &ast.BinaryOperation{
		Operation: ops.Plus,
		Left:      &ast.Integer{Value: 3},
		Right:     &ast.Integer{Value: 5},
	}
	c.Compile(
		&ast.Program{
			Statements: []ast.Statement{
				&ast.DeclarationStatement{
					Name:       "a",
					Expression: binOp,
					Type:       &ctypes.Integer{BitSize: 64},
				},
				&ast.ExpressionStatement{
					Expression: &ast.BuiltinCall{
						Name: "println",
						Parameters: []ast.Expression{
							&ast.Identifier{Name: "a"},
						},
					},
				},
			},
		},
	)
	a.AssertEqual(fmt.Sprintf("%d", 3+5), string(a.UnwrapBytes(c.Execute())))
}
