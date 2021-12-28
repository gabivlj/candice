package token

type TypeToken string

// Token is the token that we are going to parse
type Token struct {
	Type     TypeToken
	Literal  string
	Line     uint32
	Position uint32
}

const (
	// ILLEGAL

	ILLEGAL = TypeToken("ILLEGAL") // Unknown token
	EOF     = TypeToken("EOF")     // End of file

	// Identifier + Literals

	IDENT = TypeToken("IDENT") // VARIABLE NAME
	INT   = TypeToken("INT")   // 12345
	FLOAT = TypeToken("FLOAT") // 1.0

	// Operators

	AT       = TypeToken("@")
	OR       = TypeToken("||")
	AND      = TypeToken("&&")
	ANDBIN   = TypeToken("&")
	ORBIN    = TypeToken("|")
	XORBIN   = TypeToken("^")
	COMMA    = TypeToken(",")
	PLUS     = TypeToken("+")
	MINUS    = TypeToken("-")
	BANG     = TypeToken("!")
	ASTERISK = TypeToken("*")
	SLASH    = TypeToken("/")
	LT       = TypeToken("<")
	GT       = TypeToken(">")
	EQ       = TypeToken("==")
	NOTEQ    = TypeToken("!=")
	LTE      = TypeToken("<=")
	GTE      = TypeToken(">=")

	DOT = TypeToken(".")
	// Delimiters

	LPAREN = TypeToken("(")
	RPAREN = TypeToken(")")
	LBRACE = TypeToken("{")
	RBRACE = TypeToken("}")

	SEMICOLON = TypeToken(";")
	COLON     = TypeToken(":")
	ASSIGN    = TypeToken("=")

	// Keywords

	STRUCT   = TypeToken("STRUCT")
	FUNCTION = TypeToken("FUNCTION")
	LET      = TypeToken("LET")
	TRUE     = TypeToken("TRUE")
	FALSE    = TypeToken("FALSE")
	IF       = TypeToken("IF")
	ELSE     = TypeToken("ELSE")
	RETURN   = TypeToken("RETURN")

	STRING   = TypeToken("STRING")
	LBRACKET = TypeToken("[")
	RBRACKET = TypeToken("]")

	IMPORT = TypeToken("IMPORT")
	FOR    = TypeToken("FOR")

	VOID   = TypeToken("VOID")
	BREAK  = TypeToken("BREAK")
	EXTERN = TypeToken("EXTERN")
)

var keywords = map[string]TypeToken{
	"func":   FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"return": RETURN,
	"void":   VOID,
	"struct": STRUCT,
	"import": IMPORT,
	"break":  BREAK,
	"extern": EXTERN,
}

// LookupIdent Looks up in the keywords table if its a keyword, if its not it will return IDENT as a TypeToken
func LookupIdent(ident string) TypeToken {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
