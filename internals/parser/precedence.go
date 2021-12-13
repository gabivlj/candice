package parser

import (
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
)

var precedences = map[token.TypeToken]int{
	token.ASSIGN:   1,
	token.AND:      2,
	token.OR:       2,
	token.EQ:       3,
	token.NOTEQ:    3,
	token.GTE:      3,
	token.GT:       3,
	token.LT:       3,
	token.LTE:      3,
	token.PLUS:     4,
	token.MINUS:    4,
	token.XORBIN:   5,
	token.ANDBIN:   5,
	token.ORBIN:    5,
	token.BANG:     5,
	token.SLASH:    6,
	token.ASTERISK: 6,
	token.LPAREN:   8,
	token.DOT:      10,
	token.LBRACKET: 11,
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
	token.ASSIGN:   ops.TempAssign,
}

func (p *Parser) currentTokenToOperation() ops.Operation {
	return operations[p.currentToken.Type]
}
