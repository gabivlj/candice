package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gabivlj/candice/internals/ast"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/node"
	"github.com/gabivlj/candice/internals/token"
)

type prefixFunc = func() ast.Expression
type infixFunc = func(expression ast.Expression) ast.Expression

type Parser struct {
	currentToken token.Token
	peekToken    token.Token
	lexer        *lexer.Lexer

	prefixFunc map[token.TypeToken]prefixFunc
	infixFunc  map[token.TypeToken]infixFunc
	Errors     []error
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

func (p *Parser) expect(expected token.TypeToken) {
	if p.currentToken.Type != expected {
		p.addErrorMessage(fmt.Sprintf("expected: '%s', got: '%s'", string(expected), p.currentToken.Literal))
	}
}

func (p *Parser) addErrorMessage(message string) {
	errMsg := errors.New(fmt.Sprintf("expect on %d:%d 'token: %s %s': %s",
		p.currentToken.Line, p.currentToken.Position, p.currentToken.Literal, p.currentToken.Type, message))
	p.Errors = append(p.Errors, errMsg)
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:      l,
		prefixFunc: map[token.TypeToken]prefixFunc{},
		infixFunc:  map[token.TypeToken]infixFunc{},
		Errors:     []error{},
	}
	p.initBuiltinFunctions()
	p.nextToken()
	p.nextToken()
	p.registerInfixHandler(token.PLUS, p.parseInfix)
	p.registerInfixHandler(token.ASTERISK, p.parseInfix)
	p.registerInfixHandler(token.AND, p.parseInfix)
	p.registerInfixHandler(token.ANDBIN, p.parseInfix)
	p.registerInfixHandler(token.ORBIN, p.parseInfix)
	p.registerInfixHandler(token.EQ, p.parseInfix)
	p.registerInfixHandler(token.NOTEQ, p.parseInfix)
	p.registerInfixHandler(token.XORBIN, p.parseInfix)
	p.registerInfixHandler(token.GT, p.parseInfix)
	p.registerInfixHandler(token.GTE, p.parseInfix)
	p.registerInfixHandler(token.LT, p.parseInfix)
	p.registerInfixHandler(token.LTE, p.parseInfix)
	p.registerInfixHandler(token.DOT, p.parseInfix)
	p.registerInfixHandler(token.MINUS, p.parseInfix)
	p.registerInfixHandler(token.OR, p.parseInfix)
	p.registerInfixHandler(token.SLASH, p.parseInfix)
	p.registerInfixHandler(token.EQ, p.parseInfix)
	p.registerInfixHandler(token.ASSIGN, p.parseInfix)
	p.registerInfixHandler(token.LBRACKET, p.parseIndex)
	p.registerInfixHandler(token.LPAREN, p.parseCall)

	p.registerPrefixHandler(token.STRING, p.parseString)
	p.registerPrefixHandler(token.BANG, p.parsePrefixExpression)
	p.registerPrefixHandler(token.ANDBIN, p.parsePrefixExpression)
	p.registerPrefixHandler(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixHandler(token.PLUS, p.parsePrefixExpression)
	p.registerPrefixHandler(token.ASTERISK, p.parsePrefixExpression)
	p.registerPrefixHandler(token.AT, p.parseAt)
	p.registerPrefixHandler(token.IDENT, p.parseIdentifierExpression)
	p.registerPrefixHandler(token.INT, p.parseInteger)
	p.registerPrefixHandler(token.FLOAT, p.parseFloat)
	p.registerPrefixHandler(token.LPAREN, p.parseParenthesisPrefix)
	p.registerPrefixHandler(token.LBRACKET, p.parseStaticArray)
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
	defer func() {
		p.skipSemicolon()
	}()
	switch p.currentToken.Type {
	case token.EXTERN:
		return p.parseExtern()
	case token.IDENT:
		return p.parseIdentifierStatement()
	case token.ASTERISK:
		return p.parsePossibleAssignment()
	case token.IF:
		return p.parseIf()
	case token.FOR:
		return p.parseFor()
	case token.STRUCT:
		return p.parseStruct()
	case token.FUNCTION:
		return p.parseFunctionDeclaration()
	case token.RETURN:
		return p.parseReturn()
	case token.IMPORT:
		return p.parseImport()
	case token.BREAK:
		return &ast.BreakStatement{Token: p.nextToken()}
	case token.CONTINUE:
		return &ast.ContinueStatement{Token: p.nextToken()}
	default:
		{
			return p.parseExpressionStatement()
		}
	}
}

func (p *Parser) parseStaticArray() ast.Expression {
	t := p.parseType()
	if _, ok := t.(*ctypes.Array); !ok {
		p.addErrorMessage("expected array type, got: " + t.String())
	}
	p.expect(token.LBRACE)
	l := p.nextToken()
	var expressions []ast.Expression
	for p.currentToken.Type != token.RBRACE && p.currentToken.Type != token.EOF {
		if len(expressions) > 0 {
			p.expect(token.COMMA)
			p.nextToken()
		}
		expressions = append(expressions, p.parseExpression(0))
	}
	p.expect(token.RBRACE)
	p.nextToken()
	return &ast.ArrayLiteral{
		Node: &node.Node{
			Type:  t,
			Token: l,
		},
		Values: expressions,
	}
}

func (p *Parser) parseImport() ast.Statement {
	imp := p.nextToken()
	p.expect(token.IDENT)
	identifier := p.nextToken()
	var types []ctypes.Type
	for p.peekToken.Type != token.STRING && p.currentToken.Type != token.EOF {
		p.expect(token.COMMA)
		p.nextToken()
		t := p.parseType()
		if t == nil {
			p.addErrorMessage("couldn't parse type")
		}
		types = append(types, t)
	}
	p.expect(token.COMMA)
	p.nextToken()
	p.expect(token.STRING)
	path, ok := p.parseString().(*ast.StringLiteral)
	if !ok {
		p.addErrorMessage("couldn't parse string literal in import statement")
		return &ast.ImportStatement{}
	}

	return &ast.ImportStatement{
		Name:  identifier.Literal,
		Types: types,
		Path:  path,
		Token: imp,
	}
}

func (p *Parser) skipSemicolon() {
	for p.currentToken.Type == token.SEMICOLON {
		p.nextToken()
	}
}

func (p *Parser) parseReturn() ast.Statement {
	ret := p.nextToken()
	var exp ast.Expression
	if p.currentToken.Type != token.SEMICOLON {
		exp = p.parseExpression(0)
	}

	return &ast.ReturnStatement{
		Token:      ret,
		Expression: exp,
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Token:      p.currentToken,
		Expression: p.parseExpression(0),
	}
}

func (p *Parser) parseIdentifierExpression() ast.Expression {
	identifier := p.nextToken()
	return &ast.Identifier{
		Node: &node.Node{
			Type:  ctypes.TODO(),
			Token: identifier,
		},
		Name: identifier.Literal,
	}
}

func (p *Parser) parseFunctionDeclaration() ast.Statement {
	fun := p.nextToken()
	p.expect(token.IDENT)
	name := p.nextToken()
	var names []string
	var types []ctypes.Type
	p.expect(token.LPAREN)
	p.nextToken()
	for p.currentToken.Type != token.RPAREN && p.currentToken.Type != token.EOF {
		if len(names) > 0 {
			p.expect(token.COMMA)
			p.nextToken()
		}
		p.expect(token.IDENT)
		ident := p.nextToken()
		t := p.parseType()
		names = append(names, ident.Literal)
		types = append(types, t)
	}

	p.expect(token.RPAREN)
	p.nextToken()
	var returnType ctypes.Type
	if p.currentToken.Type != token.LBRACE {
		returnType = p.parseType()
	}
	block := p.parseBlock()
	return &ast.FunctionDeclarationStatement{
		Token: fun,
		FunctionType: &ctypes.Function{
			Name:       name.Literal,
			Parameters: types,
			Names:      names,
			Return:     returnType,
		},
		Block: block,
	}
}

func (p *Parser) parseStructLiteral() ast.Expression {
	literal := p.nextToken()
	p.expect(token.LBRACE)
	p.nextToken()
	var structValues []ast.StructValue

	for p.currentToken.Type != token.RBRACE && p.currentToken.Type != token.EOF {
		p.expect(token.IDENT)
		identifier := p.nextToken()
		p.expect(token.COLON)
		p.nextToken()
		expr := p.parseExpression(0)
		if len(structValues) >= 1 && p.currentToken.Type != token.RBRACE {
			p.expect(token.COMMA)
		}

		structValues = append(structValues, ast.StructValue{
			Name:       identifier.Literal,
			Expression: expr,
		})
		if p.currentToken.Type == token.COMMA {
			p.nextToken()
		}

	}

	p.expect(token.RBRACE)
	p.nextToken()
	return &ast.StructLiteral{
		Node: &node.Node{
			Type:  &ctypes.Anonymous{Name: literal.Literal},
			Token: literal,
		},
		Name:   literal.Literal,
		Values: structValues,
	}
}

func (p *Parser) parseStruct() ast.Statement {
	_ = p.nextToken()
	p.expect(token.IDENT)
	identifier := p.nextToken()
	p.expect(token.LBRACE)
	p.nextToken()
	var types []ctypes.Type
	var names []string

	for p.currentToken.Type != token.RBRACE && p.currentToken.Type != token.EOF {
		p.expect(token.IDENT)
		name := p.nextToken()
		names = append(names, name.Literal)
		t := p.parseType()
		types = append(types, t)

	}

	p.expect(token.RBRACE)
	p.nextToken()
	s := ast.StructStatement{
		Token: identifier,
		Type: &ctypes.Struct{
			Fields: types,
			Names:  names,
			Name:   identifier.Literal,
		},
	}

	return &s
}

func (p *Parser) parseIdentifierStatement() ast.Statement {
	if p.peekToken.Type == token.COLON {
		return p.parseDeclaration()
	}

	return p.parsePossibleAssignment()
}

func (p *Parser) parseDeclaration() ast.Statement {
	id := p.currentToken
	// pass id
	p.nextToken()
	// check curr == colon
	p.expect(token.COLON)
	// pass colon
	p.nextToken()
	var t = ctypes.TODO()

	if p.currentToken.Type != token.ASSIGN {
		t = p.parseType()
	}

	p.expect(token.ASSIGN)
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
		modules := []string{t.Literal}
		for p.currentToken.Type == token.DOT {
			p.nextToken()
			p.expect(token.IDENT)
			identifier := p.nextToken()
			modules = append(modules, identifier.Literal)
		}
		if len(modules) > 1 {
			return &ctypes.Anonymous{
				Modules: modules[:len(modules)-1],
				Name:    modules[len(modules)-1],
			}
		}

		return ctypes.LiteralToType(modules[0])
	}

	if p.currentToken.Type == token.FUNCTION {
		_ = p.nextToken()
		name := ""
		if p.currentToken.Type == token.IDENT {
			name = p.nextToken().Literal
		}
		p.expect(token.LPAREN)
		p.nextToken()
		var parameters []ctypes.Type
		for p.currentToken.Type != token.RPAREN && p.currentToken.Type != token.EOF {
			if len(parameters) >= 1 {
				p.expect(token.COMMA)
				p.nextToken()
			}
			parameters = append(parameters, p.parseType())
		}
		p.expect(token.RPAREN)
		p.nextToken()
		var returnType ctypes.Type
		if p.currentToken.Type == token.IDENT {
			returnType = p.parseType()
		} else if p.currentToken.Type == token.ASTERISK {
			returnType = p.parseType()
		} else if p.currentToken.Type == token.LBRACKET {
			returnType = p.parseType()
		}
		return &ctypes.Function{
			Name:       name,
			Parameters: parameters,
			Names:      []string{},
			Return:     returnType,
		}
	}

	if p.currentToken.Type == token.LBRACKET {
		p.nextToken()
		p.expect(token.INT)
		integer, err := strconv.ParseInt(p.currentToken.Literal, 10, 32)
		p.nextToken()
		if err != nil {
			p.addErrorMessage(fmt.Sprintf("couldn't parse array size because of the following expect: %s", err.Error()))
		}
		p.expect(token.RBRACKET)
		p.nextToken()
		return &ctypes.Array{Length: integer, Inner: p.parseType()}
	}

	return nil
}

func (p *Parser) parsePrefix() ast.Expression {
	fn, ok := p.prefixFunc[p.currentToken.Type]

	if !ok {
		p.addErrorMessage("unknown token to parse on prefix")
		p.nextToken()
		return &ast.Integer{Value: 1}
	}

	return fn()
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	op := p.currentTokenToOperation()
	tok := p.nextToken()
	left := &ast.PrefixOperation{
		Node: &node.Node{
			Token: tok,
			Type:  ctypes.TODO(),
		},
		Right:     p.parseExpression(p.precedencePrefix()),
		Operation: op,
	}
	return left
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
	nextPrecedence := p.precedence()
	operation := p.currentTokenToOperation()
	currentToken := p.nextToken()
	return &ast.BinaryOperation{
		Node: &node.Node{
			Type:  ctypes.TODO(),
			Token: currentToken,
		},
		Left:      expression,
		Right:     p.parseExpression(nextPrecedence),
		Operation: operation,
	}
}

func (p *Parser) parseCall(expression ast.Expression) ast.Expression {
	paren := p.nextToken()
	var expressions []ast.Expression
	for p.currentToken.Type != token.RPAREN && p.currentToken.Type != token.EOF {
		if len(expressions) > 0 {
			p.expect(token.COMMA)
			p.nextToken()
		}
		expressions = append(expressions, p.parseExpression(0))
	}
	p.expect(token.RPAREN)
	p.nextToken()
	return &ast.Call{
		Node: &node.Node{
			Type:  ctypes.TODO(),
			Token: paren,
		},
		Left:       expression,
		Parameters: expressions,
	}
}

func (p *Parser) parseIndex(expression ast.Expression) ast.Expression {
	currentToken := p.nextToken()
	i := &ast.IndexAccess{
		Node: &node.Node{
			Type:  ctypes.TODO(),
			Token: currentToken,
		},
		Left:   expression,
		Access: p.parseExpression(0),
	}
	p.expect(token.RBRACKET)
	p.nextToken()
	return i
}

func (p *Parser) parseInteger() ast.Expression {
	t := p.nextToken()
	integer, _ := strconv.ParseInt(t.Literal, 10, 64)
	return &ast.Integer{
		Node: &node.Node{
			Type:  ctypes.I32,
			Token: t,
		},
		Value: integer,
	}
}

func (p *Parser) parseFloat() ast.Expression {
	t := p.nextToken()
	float, err := strconv.ParseFloat(t.Literal, 64)
	if err != nil {
		p.addErrorMessage("couldn't parse float " + t.Literal + ", because of given error: " + err.Error())
	}
	return &ast.Float{
		Node: &node.Node{
			Type:  ctypes.F32,
			Token: t,
		},
		Value: float,
	}
}

func (p *Parser) parseString() ast.Expression {
	str := p.nextToken()
	return &ast.StringLiteral{
		Node: &node.Node{
			Type:  &ctypes.Pointer{Inner: &ctypes.Integer{BitSize: 8}},
			Token: str,
		},
		Value: str.Literal,
	}
}

func (p *Parser) parseParenthesisPrefix() ast.Expression {
	_ = p.nextToken()
	exp := p.parseExpression(0)
	p.expect(token.RPAREN)
	p.nextToken()
	return exp
}

func (p *Parser) parseAt() ast.Expression {
	p.expect(token.AT)
	at := p.nextToken()
	p.expect(token.IDENT)
	if p.peekToken.Type == token.LBRACE {
		return p.parseStructLiteral()
	}

	identifier := p.nextToken()
	p.expect(token.LPAREN)
	p.nextToken()
	functionRequirements := p.getBuiltinFunctionRequirements(identifier.Literal)
	types := p.parseBuiltinCallTypes(functionRequirements)
	expressions := p.parseBuiltinCallParameters(functionRequirements)
	p.expect(token.RPAREN)
	p.nextToken()
	return &ast.BuiltinCall{
		Node:           &node.Node{Token: at, Type: ctypes.TODO()},
		Name:           identifier.Literal,
		TypeParameters: types,
		Parameters:     expressions,
	}
}

func (p *Parser) parseBlock() *ast.Block {
	if p.currentToken.Type == token.LBRACE {
		p.nextToken()
		block := &ast.Block{
			Statements: []ast.Statement{},
		}
		for p.currentToken.Type != token.RBRACE && p.currentToken.Type != token.EOF {
			block.Statements = append(block.Statements, p.parseStatement())
		}
		p.expect(token.RBRACE)
		p.nextToken()
		return block
	}

	stmt := p.parseStatement()
	return &ast.Block{Statements: []ast.Statement{stmt}}
}

func (p *Parser) parseIf() ast.Statement {
	ifToken := p.nextToken()
	condition := p.parseExpression(0)
	block := p.parseBlock()
	var elseWithConditions []*ast.ConditionPlusBlock
	var elseBlock *ast.Block
	for p.currentToken.Type == token.ELSE {
		p.nextToken()
		if p.currentToken.Type != token.IF {
			elseBlock = p.parseBlock()
			break
		} else if p.currentToken.Type == token.IF {
			p.nextToken()
			expression := p.parseExpression(0)
			block := p.parseBlock()
			elseWithConditions = append(elseWithConditions, &ast.ConditionPlusBlock{
				Block:     block,
				Condition: expression,
			})
		}
	}

	return &ast.IfStatement{
		Token:     ifToken,
		Condition: condition,
		Block:     block,
		ElseIfs:   elseWithConditions,
		Else:      elseBlock,
	}
}

func (p *Parser) parseBuiltinCallTypes(builtinRequirements BuiltinFunctionParseRequirements) []ctypes.Type {
	var types []ctypes.Type
	if builtinRequirements.Types == 0 {
		return types
	}

	for i := 0; i < builtinRequirements.Types; i++ {
		types = append(types, p.parseType())
		if i+1 < builtinRequirements.Types || builtinRequirements.Parameters > 0 {
			if p.currentToken.Type == token.RPAREN {
				p.addErrorMessage(fmt.Sprintf("expected %d types, got=%d", builtinRequirements.Types, i))
				break
			}

			p.expect(token.COMMA)
			p.nextToken()
		} else {
			if builtinRequirements.Types > 0 {
				p.expect(token.COMMA)
				p.nextToken()
			}
		}
	}

	return types
}

func (p *Parser) parseBuiltinCallParameters(builtinRequirements BuiltinFunctionParseRequirements) []ast.Expression {
	var expressions []ast.Expression
	if builtinRequirements.Parameters == 0 {
		return expressions
	}

	if builtinRequirements.Parameters == UndefinedNumberOfParameters {
		for {
			expressions = append(expressions, p.parseExpression(0))
			if p.currentToken.Type == token.RPAREN {
				return expressions
			}
			p.expect(token.COMMA)
			if p.currentToken.Type != token.COMMA {
				return expressions
			}
		}
	}

	for i := 0; i < builtinRequirements.Parameters; i++ {
		expressions = append(expressions, p.parseExpression(0))
		if i+1 < builtinRequirements.Parameters {
			if p.currentToken.Type == token.RPAREN {
				p.addErrorMessage(fmt.Sprintf("expected %d expressions, got=%d", builtinRequirements.Types, i))
				break
			}

			p.expect(token.COMMA)
			p.nextToken()
		}

	}

	return expressions
}

func (p *Parser) parseFor() ast.Statement {
	forToken := p.nextToken()
	if p.currentToken.Type == token.LBRACE {
		return &ast.ForStatement{
			Token:                forToken,
			Condition:            nil,
			InitializerStatement: nil,
			Operation:            nil,
			Block:                p.parseBlock(),
		}
	}

	possibleStatement := p.parseStatement()
	var assignment, operation ast.Statement
	var condition ast.Expression
	_, isAssignment := possibleStatement.(*ast.AssignmentStatement)
	_, isDeclaration := possibleStatement.(*ast.DeclarationStatement)
	if isDeclaration || isAssignment {
		assignment = possibleStatement
	} else {
		possibleExpression, isExpression := possibleStatement.(*ast.ExpressionStatement)

		// for <statement>
		if !isExpression {
			return &ast.ForStatement{
				Token:                forToken,
				Condition:            nil,
				InitializerStatement: nil,
				Operation:            nil,
				Block:                &ast.Block{Statements: []ast.Statement{possibleStatement}},
			}
		}

		// for <statement>
		if p.currentToken.Type != token.LBRACE {
			return &ast.ForStatement{
				Token:                forToken,
				Condition:            nil,
				InitializerStatement: nil,
				Operation:            nil,
				Block:                &ast.Block{Statements: []ast.Statement{possibleStatement}},
			}
		}

		// for <condition> { /*block*/ }

		condition = possibleExpression.Expression
		block := p.parseBlock()
		return &ast.ForStatement{
			Token:                forToken,
			Condition:            condition,
			InitializerStatement: nil,
			Operation:            nil,
			Block:                block,
		}
	}
	condition = p.parseExpression(0)
	p.expect(token.SEMICOLON)
	p.nextToken()
	if p.currentToken.Type != token.LBRACE {
		operation = p.parsePossibleAssignment()
	}
	block := p.parseBlock()
	return &ast.ForStatement{
		Token:                forToken,
		Condition:            condition,
		InitializerStatement: assignment,
		Operation:            operation,
		Block:                block,
	}
}

func (p *Parser) parseExtern() ast.Statement {
	extern := p.nextToken()
	p.expect(token.FUNCTION)
	t := p.parseType()
	if _, ok := t.(*ctypes.Function); !ok {
		p.addErrorMessage("expected external function, got " + t.String())
		return &ast.ExternStatement{}
	}

	if fun, ok := t.(*ctypes.Function); !ok || fun.Name == "" {
		p.addErrorMessage("badly formed external function")
	}

	return &ast.ExternStatement{
		Token: extern,
		Type:  t,
	}
}
