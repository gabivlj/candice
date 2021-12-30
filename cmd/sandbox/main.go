package main

import (
	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"github.com/gabivlj/candice/pkg/a"
	"os"
	"strings"
)

func main() {
	code, _ := os.ReadFile("./sandbox.cd")

	l := lexer.New(string(code))
	p := parser.New(l)
	s := semantic.New()
	tree := p.Parse()
	s.Analyze(tree)
	c := compiler.New()
	c.Compile(tree)
	bytes, err := c.Execute()
	a.AssertErr(err)
	a.Assert(strings.TrimSpace(string(bytes)) == "3")
}
