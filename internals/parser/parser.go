package parser

import (
	"errors"
	"fmt"
	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/node"
	"github.com/gabivlj/candice/internals/ops"
	"github.com/gabivlj/candice/internals/token"
	"strconv"
)

type prefixFunc = func() ast.Expression
type infixFunc = func(expression ast.Expression) ast.Expression

type Parser struct {
	currentToken token.Token
	peekToken    token.Token
	lexer        *lexer.Lexer

	prefixFunc map[token.TypeToken]prefixFunc
	infixFunc  map[token.TypeToken]infixFunc
	errors     []error
}

func (p *Parser) registerPrefixHandler(tokenType token.TypeToken, prefixFunc prefixFunc) {
	p.prefixFunc[tokenType] = prefixFunc
}

func (p *Parser) registerInfixHandler(tokenType token.TypeToken, infixFunc infixFunc) {
	p.infixFunc[tokenType] = infixFunc
}

func (p *Parser) nextToken() token.Token {
	prev := p.currentToken
	p.currentToken, p.peekToken = p.peekToken, p.lexer.NextToken()
	return prev
}

func (p *Parser) error(expected token.TypeToken) {
	if p.currentToken.Type != expected {
		p.addErrorMessage(fmt.Sprintf("expected: '%s', got: '%s'", string(expected), p.currentToken.Literal))
	}
}

func (p *Parser) addErrorMessage(message string) {
	errMsg := errors.New(fmt.Sprintf("error on %d-%d\n - %s", p.currentToken.Line, p.currentToken.Position, message))
	p.errors = append(p.errors, errMsg)
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:      l,
		prefixFunc: map[token.TypeToken]prefixFunc{},
		infixFunc:  map[token.TypeToken]infixFunc{},
		errors:     []error{},
	}
	p.nextToken()
	p.nextToken()
	p.registerInfixHandler(token.PLUS, p.parseInfix)

	p.registerPrefixHandler(token.INT, func() ast.Expression {
		t := p.nextToken()
		integer, _ := strconv.ParseInt(t.Literal, 10, 64)
		return &ast.Integer{
			Node: &node.Node{
				Type:  &ctypes.Integer{BitSize: 64},
				Token: t,
			},
			Value: integer,
		}
	})

	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}
	for p.currentToken.Type != token.EOF {
		program.Statements = append(program.Statements, p.parseStatement())
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.IDENT:
		{
			return p.parseIdentifierStatement()
		}
	default:
		{
			return p.parseExpressionStatement()
		}
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Token:      p.currentToken,
		Expression: p.parseExpression(0),
	}
}

func (p *Parser) parseIdentifierStatement() ast.Statement {
	if p.peekToken.Type == token.COLON {
		return p.parseDeclaration()
	}
	return nil
}

func (p *Parser) parseDeclaration() ast.Statement {
	id := p.currentToken
	// pass id
	p.nextToken()
	// check curr == colon
	p.error(token.COLON)
	// pass colon
	p.nextToken()

	var t ctypes.Type = &ctypes.Anonymous{Name: "<INFER>"}

	if p.currentToken.Type != token.ASSIGN {
		t = p.parseType()
	}

	p.error(token.ASSIGN)
	// pass assign
	p.nextToken()

	return &ast.DeclarationStatement{
		Token:      id,
		Name:       id.Literal,
		Type:       t,
		Expression: p.parseExpression(0),
	}
}

func (p *Parser) parseType() ctypes.Type {
	if p.currentToken.Type == token.ASTERISK {
		p.nextToken()
		return &ctypes.Pointer{Inner: p.parseType()}
	}

	if p.currentToken.Type == token.IDENT {
		t := p.nextToken()
		return ctypes.LiteralToType(t.Literal)
	}

	if p.currentToken.Type == token.LBRACKET {
		p.nextToken()
		p.error(token.INT)
		integer, err := strconv.ParseInt(p.currentToken.Literal, 10, 32)
		p.nextToken()
		if err != nil {
			p.addErrorMessage(fmt.Sprintf("couldn't parse array size because of the following error: %s", err.Error()))
		}
		p.error(token.RBRACKET)
		p.nextToken()
		return &ctypes.Array{Length: integer, Inner: p.parseType()}
	}

	if p.currentToken.Type == token.VOID {
		return &ctypes.Void{}
	}

	return nil
}

func (p *Parser) parsePrefix() ast.Expression {
	fn, ok := p.prefixFunc[p.currentToken.Type]

	if !ok {
		p.addErrorMessage("unknown token to parse")
		p.nextToken()
		return &ast.Integer{Value: 1}
	}

	return fn()
}

func (p *Parser) parseExpression(prec int) ast.Expression {
	prefixExpr := p.parsePrefix()
	for p.currentToken.Type != token.SEMICOLON && p.currentToken.Type != token.EOF && prec < p.precedence() {
		infix := p.infixFunc[p.currentToken.Type]
		if infix == nil {
			return prefixExpr
		}
		prefixExpr = infix(prefixExpr)
	}
	return prefixExpr
}

func (p *Parser) parseInfix(expression ast.Expression) ast.Expression {
	nextPrec := p.precedence()
	token := p.nextToken()
	return &ast.BinaryOperation{
		Node: &node.Node{
			Type:  nil,
			Token: token,
		},
		Left:      expression,
		Right:     p.parseExpression(nextPrec),
		Operation: ops.Plus,
	}
}
