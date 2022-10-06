package token

type TypeToken string

// Token is the token that we are going to parse
type Token struct {
	Type     TypeToken
	Literal  string
	Line     uint32
	Position uint32
	// This is only for retrieving error messages, maybe we shouldn't increase 4 bytes our tokens because of
	// that.
	OverallPosition int
}

const (
	ILLEGAL = TypeToken("ILLEGAL") // Unknown token
	EOF     = TypeToken("EOF")     // End of file

	// Identifier + Literals
	IDENT  = TypeToken("IDENT")  // VARIABLE NAME
	INT    = TypeToken("INT")    // 12345
	FLOAT  = TypeToken("FLOAT")  // 1.0
	HEX    = TypeToken("HEX")    // 0x12191
	BINARY = TypeToken("BINARY") // 0b1101010
	STRING = TypeToken("STRING")

	// Operators
	AT           = TypeToken("@")
	OR           = TypeToken("||")
	AND          = TypeToken("&&")
	ANDBIN       = TypeToken("&")
	ORBIN        = TypeToken("|")
	XORBIN       = TypeToken("^")
	COMMA        = TypeToken(",")
	PLUS         = TypeToken("+")
	MINUS        = TypeToken("-")
	BANG         = TypeToken("!")
	ASTERISK     = TypeToken("*")
	SLASH        = TypeToken("/")
	LT           = TypeToken("<")
	LS           = TypeToken("<<")
	RS           = TypeToken(">>")
	GT           = TypeToken(">")
	EQ           = TypeToken("==")
	NOTEQ        = TypeToken("!=")
	LTE          = TypeToken("<=")
	GTE          = TypeToken(">=")
	DOT          = TypeToken(".")
	DOUBLE_DOT   = TypeToken("..")
	DOUBLE_PLUS  = TypeToken("++")
	DOUBLE_MINUS = TypeToken("--")
	MODULO       = TypeToken("%")
	ARROW        = TypeToken("->")

	// Delimiters
	LPAREN    = TypeToken("(")
	RPAREN    = TypeToken(")")
	LBRACE    = TypeToken("{")
	RBRACE    = TypeToken("}")
	LBRACKET  = TypeToken("[")
	RBRACKET  = TypeToken("]")
	SEMICOLON = TypeToken(";")
	COLON     = TypeToken(":")
	ASSIGN    = TypeToken("=")

	// Keywords
	TYPE     = TypeToken("type")
	PUBLIC   = TypeToken("pub")
	UNION    = TypeToken("union")
	STRUCT   = TypeToken("STRUCT")
	FUNCTION = TypeToken("FUNCTION")
	TRUE     = TypeToken("TRUE")
	FALSE    = TypeToken("FALSE")
	IF       = TypeToken("IF")
	ELSE     = TypeToken("ELSE")
	RETURN   = TypeToken("RETURN")
	IMPORT   = TypeToken("IMPORT")
	FOR      = TypeToken("FOR")
	BREAK    = TypeToken("BREAK")
	CONTINUE = TypeToken("CONTINUE")
	EXTERN   = TypeToken("EXTERN")
	AS       = TypeToken("AS")
	MACRO_IF = TypeToken("#IF")

	SWITCH  = TypeToken("SWITCH")
	CASE    = TypeToken("CASE")
	DEFAULT = TypeToken("DEFAULT")
	CONST   = TypeToken("CONSTANT")
	CHAR    = TypeToken("CHAR")
)

var keywords = map[string]TypeToken{
	"func":     FUNCTION,
	"pub":      PUBLIC,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"return":   RETURN,
	"struct":   STRUCT,
	"import":   IMPORT,
	"break":    BREAK,
	"extern":   EXTERN,
	"continue": CONTINUE,
	"type":     TYPE,
	"as":       AS,
	"union":    UNION,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"const":    CONST,
}

// LookupIdent Looks up in the keywords table if its a keyword, if its not it will return IDENT as a TypeToken
func LookupIdent(ident string) TypeToken {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
