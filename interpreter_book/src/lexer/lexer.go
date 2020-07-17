package lexer

import (
	"github.com/Asalle/monkey-interpreter/token"
)

type Lexer struct {
	input string
}

func MakeLexer(input string) Lexer {
	return Lexer{input: input}
}

func (l Lexer) NextToken() token.Token {
	return token.Token{}

}
