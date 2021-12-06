package main

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/ops"
)

func main() {
	decl := &ast.DeclarationStatement{
		Name:       "cool",
		Type:       &ctypes.Array{Inner: &ctypes.Pointer{Inner: &ctypes.Integer{BitSize: 64}}, Length: 5},
		Expression: &ast.Integer{Value: 5},
	}
	i := ast.IfStatement{
		Condition: &ast.BinaryOperation{
			Left: &ast.Integer{
				Value: 3,
			},
			Right: &ast.Integer{
				Value: 0,
			},
			Operation: ops.Multiply,
		},
		Block: &ast.Block{Statements: []ast.Statement{
			decl,
		}},
		ElseIfs: []*ast.ConditionPlusBlock{{
			Condition: &ast.Integer{Value: 3},
			Block:     &ast.Block{Statements: []ast.Statement{decl}},
		},
		},
		Else: &ast.Block{
			Statements: []ast.Statement{&ast.DeclarationStatement{
				Name:       "cool",
				Type:       &ctypes.Array{Inner: &ctypes.Pointer{Inner: &ctypes.Integer{BitSize: 64}}, Length: 5},
				Expression: &ast.Integer{Value: 5},
			}}},
	}
	fmt.Println(i.String())
}
