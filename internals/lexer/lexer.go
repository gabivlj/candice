package lexer

import (
	"github.com/gabivlj/candice/internals/token"
)

type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // next position after current char
	ch           byte // current char
	line         uint32
	column       uint32
}

// New Returns a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 1}
	// Initialize to first char.
	l.readChar()
	return l
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		if l.ch == '\n' {
			l.line++
			l.column = 1
		}
		l.readChar()
	}
}

// Returns if there is a combination with the next char, else returns otherwise param
func (l *Lexer) peekerForTwoChars(expect byte, otherwise token.Token, t token.TypeToken) token.Token {
	// Peek next character
	peek := l.peekChar()
	// Checks if the peeked character is the expected one
	if peek == expect {
		// Store the current character
		ch := l.ch
		// Next character
		l.readChar()
		// Return the token for that combination
		return token.Token{Type: t, Literal: string(ch) + string(peek)}
	}
	// Otherwise return the token that it falls to
	return otherwise
}

// NextToken Returns the next token of an input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '.':
		tok = l.newToken(token.DOT, l.ch)
	case ':':
		tok = l.newToken(token.COLON, l.ch)
	case '[':
		tok = l.newToken(token.LBRACKET, l.ch)
	case ']':
		tok = l.newToken(token.RBRACKET, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '=':
		tok = l.peekerForTwoChars('=', l.newToken(token.ASSIGN, '='), token.EQ)
	case '!':
		tok = l.peekerForTwoChars('=', l.newToken(token.BANG, '!'), token.NOTEQ)
	case ';':
		tok = l.newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = l.newToken(token.LPAREN, l.ch)
	case '+':
		tok = l.newToken(token.PLUS, l.ch)
	case '-':
		tok = l.newToken(token.MINUS, l.ch)
	case '/':
		tok = l.newToken(token.SLASH, l.ch)
	case '*':
		tok = l.newToken(token.ASTERISK, l.ch)
	case '<':
		tok = l.peekerForTwoChars('=', l.newToken(token.LT, l.ch), token.LTE)
	case '>':
		tok = l.peekerForTwoChars('=', l.newToken(token.GT, l.ch), token.GTE)
	case '&':
		tok = l.peekerForTwoChars('&', l.newToken(token.ANDBIN, l.ch), token.AND)
	case '|':
		tok = l.peekerForTwoChars('|', l.newToken(token.ORBIN, l.ch), token.OR)
	case '^':
		tok = l.newToken(token.XORBIN, l.ch)
	case ')':
		tok = l.newToken(token.RPAREN, l.ch)
	case ',':
		tok = l.newToken(token.COMMA, l.ch)
	case '{':
		tok = l.newToken(token.LBRACE, l.ch)
	case '}':
		tok = l.newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// returns the whole identifier
			tok.Literal = l.readIdentifier()
			// lookup the literal in the keyword table, if it doesn't exist it's a IDENT.
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = l.newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()

	return tok
}

func (l *Lexer) readString() string {
	s := ""
	l.readChar()
	for l.ch != '"' && l.ch != 0 {
		s += string(l.ch)
		l.readChar()
	}
	return s
}

func (l *Lexer) newToken(tokenType token.TypeToken, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: l.line, Position: l.column}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	defer func() {
		// next position is current position now
		l.position = l.readPosition
		l.column++
		// point to the next char
		l.readPosition++
	}()
	// EOF
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}
	// read next position.
	l.ch = l.input[l.readPosition]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}