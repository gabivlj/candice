package main

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
)

func main() {
	tree := parser.New(lexer.New("thing.thing2.thing3[3].thing4.thing5.thing6")).Parse()
	fmt.Printf("%v", tree.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.BinaryOperation).Left.(*ast.BinaryOperation).Right)
}
