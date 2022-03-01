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
	token.LS:       5,
	token.RS:       5,
	token.ORBIN:    5,
	token.BANG:     8,
	token.SLASH:    6,
	token.ASTERISK: 6,
	token.LPAREN:   9,
	token.LBRACKET: 11,
	token.DOT:      12,
	token.AS:       7,
	// This is only for showcasing errors.
	token.DOUBLE_MINUS: 20,
	token.DOUBLE_PLUS:  20,
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
}

func (p *Parser) currentTokenToOperation() ops.Operation {
	return operations[p.currentToken.Type]
}
