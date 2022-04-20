package parser

import (
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
)

var precedences = map[token.TypeToken]int{
	token.ASSIGN: 1,
	token.COMMA:  2,
	token.AND:    3,
	token.OR:     3,
	token.EQ:     4,
	token.NOTEQ:  4,
	token.GTE:    4,
	token.GT:     4,
	token.LT:     4,
	token.LTE:    4,
	token.PLUS:   5,
	token.MINUS:  5,
	token.XORBIN: 6,
	token.ANDBIN: 6,
	token.LS:     6,
	token.RS:     6,
	token.ORBIN:  6,

	// Used for prefix operators, not only for '!'
	token.BANG:     9,
	token.SLASH:    7,
	token.ASTERISK: 7,
	token.MODULO:   7,
	token.LPAREN:   10,
	token.LBRACKET: 12,
	token.DOT:      13,
	token.AS:       8,

	// This is only for showcasing errors.
	token.DOUBLE_MINUS: 21,
	token.DOUBLE_PLUS:  21,
}

func (p *Parser) precedencePrefix() int {
	return precedences[token.BANG]
}

func (p *Parser) precedence() int {
	return precedences[p.currentToken.Type]
}

var operations = map[token.TypeToken]ops.Operation{
	token.AND:          ops.AND,
	token.OR:           ops.OR,
	token.XORBIN:       ops.BinaryXOR,
	token.ANDBIN:       ops.BinaryAND,
	token.ORBIN:        ops.BinaryOR,
	token.PLUS:         ops.Add,
	token.SLASH:        ops.Divide,
	token.ASTERISK:     ops.Multiply,
	token.MINUS:        ops.Subtract,
	token.GT:           ops.GreaterThan,
	token.GTE:          ops.GreaterThanEqual,
	token.LT:           ops.LessThan,
	token.LTE:          ops.LessThanEqual,
	token.BANG:         ops.Bang,
	token.AT:           ops.At,
	token.EQ:           ops.Equals,
	token.NOTEQ:        ops.NotEquals,
	token.DOT:          ops.Dot,
	token.ASSIGN:       ops.TempAssign,
	token.AS:           ops.As,
	token.DOUBLE_PLUS:  ops.AddOne,
	token.DOUBLE_MINUS: ops.SubtractOne,
	token.LS:           ops.LeftShift,
	token.RS:           ops.RightShift,
	token.MODULO:       ops.Modulo,
}

func (p *Parser) currentTokenToOperation() ops.Operation {
	return operations[p.currentToken.Type]
}
