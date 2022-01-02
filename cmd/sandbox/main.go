package main

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
)

func main() {
	code := `*pointAlloc = @Point { x: @cast(i64, 43), y: @cast(i64, 55), points: @alloc(Point, 33) }`

	l := lexer.New(string(code))
	p := parser.New(l)
	//s := semantic.New()
	tree := p.Parse()
	fmt.Printf("%v", tree.Statements[0].(*ast.AssignmentStatement).Left)
	//s.Analyze(tree)
	//if len(s.Errors) > 0 {
	//	log.Println(s.Errors)
	//	return
	//}
	//c := compiler.New()
	//c.Compile(tree)
	//bytes, err := c.Execute()
	//a.AssertErr(err)
	//a.Assert(strings.TrimSpace(string(bytes)) == "-3 0 -3", strings.TrimSpace(string(bytes)))
}
