package parser

import (
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
)

var precedences = map[token.TypeToken]int{
	token.AND:      1,
	token.OR:       1,
	token.EQ:       2,
	token.NOTEQ:    2,
	token.GTE:      2,
	token.GT:       2,
	token.LT:       2,
	token.LTE:      2,
	token.PLUS:     3,
	token.MINUS:    3,
	token.XORBIN:   4,
	token.ANDBIN:   4,
	token.ORBIN:    4,
	token.BANG:     4,
	token.SLASH:    5,
	token.ASTERISK: 5,
	token.LPAREN:   7,
	token.DOT:      9,
}

func (p *Parser) precedencePrefix() int {
	return precedences[token.BANG]
}

func (p *Parser) precedence() int {
	return precedences[p.currentToken.Type]
}

var operations = map[token.TypeToken]ops.Operation{
	token.AND:      ops.AND,
	token.OR:       ops.OR,
	token.XORBIN:   ops.BinaryXOR,
	token.ANDBIN:   ops.BinaryAND,
	token.ORBIN:    ops.BinaryOR,
	token.PLUS:     ops.Plus,
	token.SLASH:    ops.Divide,
	token.ASTERISK: ops.Multiply,
	token.MINUS:    ops.Minus,
	token.GT:       ops.GreaterThan,
	token.GTE:      ops.GreaterThanEqual,
	token.LT:       ops.LessThan,
	token.LTE:      ops.LessThanEqual,
	token.BANG:     ops.Bang,
	token.AT:       ops.At,
	token.EQ:       ops.Equals,
	token.NOTEQ:    ops.NotEquals,
	token.DOT:      ops.Dot,
}

func (p *Parser) currentTokenToOperation() ops.Operation {
	return operations[p.currentToken.Type]
}
