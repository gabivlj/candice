package parser

import (
	"github.com/gabivlj/candice/internals/token"
)

var precedences = map[token.TypeToken]int{
	token.AND:      1,
	token.OR:       1,
	token.PLUS:     3,
	token.MINUS:    3,
	token.XORBIN:   3,
	token.ANDBIN:   3,
	token.ORBIN:    3,
	token.SLASH:    5,
	token.ASTERISK: 5,
	token.LPAREN:   7,
	token.DOT:      9,
}

func (p *Parser) precedence() int {
	return precedences[p.currentToken.Type]
}
