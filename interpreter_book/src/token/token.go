package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = TokenType("ILLEGAL")
	EOF     = TokenType("EOF")

	// Identifiers + literals
	IDENT  = TokenType("IDENT") // add, foobar, x, y, ...
	INT    = TokenType("INT")
	ASSIGN = TokenType("=")
	PLUS   = TokenType("+")

	// Delimiters
	COMMA     = TokenType(",")
	SEMICOLON = TokenType(";")
	LPAREN    = TokenType("(")
	RPAREN    = TokenType(")")
	LBRACE    = TokenType("{")
	RBRACE    = TokenType("}")

	// Keywords
	FUNCTION = TokenType("FUNCTION")
	LET      = TokenType("LET")
)
