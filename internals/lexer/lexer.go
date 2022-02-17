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
	l := &Lexer{input: input, line: 1, column: 0}
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

func (l *Lexer) skipUntilJL() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.readChar()
	l.line++
	l.column = 1
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
		return token.Token{Type: t, Literal: string(ch) + string(peek), Line: l.line, Position: l.column}
	}
	// Otherwise return the token that it falls to
	return otherwise
}

func (l *Lexer) getMacroToken() token.Token {
	macro := l.ch
	l.readChar()
	identifier := l.readIdentifier()
	switch identifier {
	case "if":

		return l.newToken(token.MACRO_IF, macro)
	}

	return l.newToken(token.ILLEGAL, macro)
}

// NextToken Returns the next token of an input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	for l.ch == '/' && l.peekChar() == '/' {
		l.skipUntilJL()
		l.skipWhiteSpace()
	}
	l.skipWhiteSpace()
	switch l.ch {
	case '#':
		tok = l.getMacroToken()
		return tok
	case '@':
		tok = l.newToken(token.AT, l.ch)
	case '.':
		tok = l.peekerForTwoChars('.', l.newToken(token.DOT, l.ch), token.DOUBLE_DOT)
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
			tok.Position = l.column
			tok.Line = l.line
			return tok
		}
		if isDigit(l.ch) {
			tok = l.parseNumericToken()
			return tok
		}
		tok = l.newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()

	return tok
}

func (l *Lexer) parseNumericToken() token.Token {
	var tok token.Token
	tok.Literal, tok.Type = l.readNumber()
	tok.Position = l.column
	tok.Line = l.line
	return tok
}

func (l *Lexer) readString() string {
	s := ""
	l.readChar()
	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\\' {
			// It's a jumpline
			if l.peekChar() == 'n' {
				l.readChar()
				l.readChar()
				s += string(byte(10))
				continue
			}

			// Read it as a literal
			s += string(l.ch)
			l.readChar()
			s += string(l.ch)
			l.readChar()
			continue
		}

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

func (l *Lexer) readHex() (string, token.TypeToken) {
	position := l.position
	// '0'
	l.readChar()
	// 'x'
	l.readChar()
	// Numbers and characters
	for isDigit(l.ch) || (l.ch >= 'A' && l.ch <= 'F') || (l.ch >= 'a' && l.ch <= 'f') {
		l.readChar()
	}
	return l.input[position:l.position], token.HEX
}

func (l *Lexer) readBin() (string, token.TypeToken) {
	position := l.position
	// '0'
	l.readChar()
	// 'b'
	l.readChar()
	// Numbers and characters
	for l.ch == '0' || l.ch == '1' {
		l.readChar()
	}
	return l.input[position:l.position], token.BINARY
}

func (l *Lexer) readNumber() (string, token.TypeToken) {

	position := l.position
	tokenType := token.INT

	if l.ch == '0' && l.peekChar() == 'x' {
		return l.readHex()
	}

	if l.ch == '0' && l.peekChar() == 'b' {
		return l.readBin()
	}

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' {
		l.readChar()
		tokenType = token.FLOAT
	}
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position], tokenType
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
