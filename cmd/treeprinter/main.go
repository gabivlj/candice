package main

import (
	"fmt"

	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/tree_printer"
)

func main() {
	l := lexer.New(`
	hello(4).hell(3, 3);
	struct Hell {
		e i32
		b i32
	}


	`)
	tree := parser.New(l).Parse()
	fmt.Print(tree_printer.ProcessProgram(tree))
}
