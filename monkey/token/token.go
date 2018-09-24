package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers
	IDENT = "IDENT" // variables e.g
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Comparators
	GT = ">"
	LT = "<"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"let":  LET,
	"func": FUNCTION,
}

// LookupKeywordIdent looks up to see if the input is a keyword. Returns IDENT by default
func LookupKeywordIdent(input string) TokenType {
	if tok, ok := keywords[input]; ok {
		return tok
	}
	return IDENT
}
