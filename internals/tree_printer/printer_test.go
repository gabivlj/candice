package tree_printer

import (
	"testing"

	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/pkg/a"
)

func TestSimple(t *testing.T) {
	v := ConnectString("tag", []string{"otherthing"})
	a.AssertEqual(v, `tag═otherthing`)
}

func TestMoreComplex(t *testing.T) {
	v := ConnectString("tag", []string{ConnectString("other", []string{"o", "e", "i", ConnectString("other", []string{"o", "e", "i"})})})
	a.AssertEqual(v, `tag═other╦o
         ╠e
         ╠i
         ╚other╦o
               ╠e
               ╚i`)
}

func TestParsing(t *testing.T) {
	l := lexer.New("3 + 3 - 3 * 3")
	p := parser.New(l)
	ast := p.Parse()
	s := Process(ast.Statements[0])
	a.AssertEqual(s, `ast.ExpressionStatement═ast.BinaryOperation╦ast.BinaryOperation╦3 (i32)
                                           ║                   ╠'+'
                                           ║                   ╚3 (i32)
                                           ╠'-'
                                           ╚ast.BinaryOperation╦3 (i32)
                                                               ╠'*'
                                                               ╚3 (i32)`)
}

func TestAnonymousFunction(t *testing.T) {
	l := lexer.New("a := func(a i32, b i64) i32 {};")
	p := parser.New(l)
	ast := p.Parse()
	s := Process(ast.Statements[0])
	a.AssertEqual(s, `ast.DeclarationStatement╦a
                        ╚ast.AnonymousFunction╦types.Parameters╦a i32
                                              ║                ╚b i64
                                              ╠types.Return═i32
                                              ╚`)
}
