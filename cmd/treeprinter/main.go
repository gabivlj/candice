package main

import (
	"fmt"

	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/tree_printer"
)

func main() {
	l := lexer.New(`
3 * 3 + a.a * 4 + 4;



	`)
	tree := parser.New(l).Parse()
	for _, statement := range tree.Statements {
		fmt.Println(tree_printer.Process(statement))
	}
}
