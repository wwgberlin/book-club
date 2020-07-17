package lexer

import (
	"github.com/Asalle/monkey-interpreter/token"
	"testing"
)

func Test_NextToken(t *testing.T) {
	input := "=+{}();,"

	cases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
	}

	l := MakeLexer(input)

	for _, tt := range cases {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Errorf("Wrong token type: expected: %v, got: %v", tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Errorf("Wrong token literal: expected: %v, got: %v", tt.expectedLiteral, token.Literal)
		}
	}
}
