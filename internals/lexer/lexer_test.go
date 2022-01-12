package lexer

import (
	"github.com/gabivlj/candice/internals/token"
	"github.com/gabivlj/candice/pkg/a"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	l := New("{}==!==(),&&&\n|||")
	for el := l.NextToken(); el.Type != token.EOF; el = l.NextToken() {
		//log.Println(el)
	}
}

func TestLexer_Float(t *testing.T) {
	l := New("3.3333333")
	floatToken := l.NextToken()
	a.Assert(floatToken.Type == token.FLOAT)
	a.Assert(floatToken.Literal == "3.3333333")
}
