package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tok TokenType, ch byte) Token {
	return Token{Type: tok, Literal: string(ch)}
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookUpIdent(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
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
	BANG      = TokenType("!")
	MINUS     = TokenType("-")
	SLASH     = TokenType("/")
	ASTERISK  = TokenType("*")
	LT        = TokenType("<")
	GT        = TokenType(">")

	// Keywords
	FUNCTION = TokenType("FUNCTION")
	LET      = TokenType("LET")
	IF       = TokenType("if")
	ELSE     = TokenType("else")
	TRUE     = TokenType("true")
	FALSE    = TokenType("false")
	RETURN   = TokenType("return")
)
