package lexer

import (
	"github.com/gabivlj/candice/internals/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	l := New("{}==!==(),&&&\n|||")
	for el := l.NextToken(); el.Type != token.EOF; el = l.NextToken() {
		//log.Println(el)
	}
}
