package parser

import (
	"monkey/lexer"

	"strings"
	"testing"
)

func TestBlockStatementsParseError(t *testing.T) {

	tests := []string{
		"if(1){1",
		"fn(x,y){x+y",
		"let max = fn(x,y){if(x>y){x}else{y",
	}

	for _, input := range tests {
		errors := parseErrors(input)
		if len(errors) == 0 {
			t.Errorf("Parser failed to detect missing right brace(s) in: %q", input)
			continue
		}

		/* in case you want to have a look at the errors:
		t.Logf("parser has %d errors", len(p.errors))
		for _, msg := range p.errors {
			t.Errorf("parser error: %q", msg)
		}
		*/
	}
}

func TestFunctionLiteralsInvalidParameterSingle(t *testing.T) {

	for _, tt := range invalidParams {
		input := "fn(" + tt + "){}"
		errors := parseErrors(input)
		if len(errors) == 0 {
			t.Errorf("Parser failed to detect invalid Parameter in: %q ", input)
		}
	}
}

func TestFunctionLiteralsInvalidParameterOutsider(t *testing.T) {

	for _, tt := range invalidParams {
		input := "fn(a," + tt + ",b){}"
		errors := parseErrors(input)
		if len(errors) == 0 {
			t.Errorf("Parser failed to detect invalid Parameter in: %q ", input)
		}
	}
}

func TestFunctionLiteralsInvalidParametersCrowd(t *testing.T) {

	input := "fn(" + strings.Join(invalidParams, ",") + "){} @"
	input = strings.ReplaceAll(input, ",)", "") // we remove ")" as parameter
	errors := parseErrors(input)
	if len(errors) <= 1 {
		t.Errorf("Parser failed to detect a lot of errors in: %q", input)
	} else {
		lastMsg := errors[len(errors)-1] // C
		if lastMsg != "no prefix parse function for ILLEGAL found" {
			t.Errorf("Parser failed to parse the whole input for:  %q", input)
			t.Errorf("Last error: %q", lastMsg)
		}
	}
}

func parseErrors(input string) []string {

	l := lexer.New(input)
	p := New(l)
	p.ParseProgram()

	return p.Errors()
}

var invalidParams = []string{
	"@",      //ILLEGAL
	"0",      //INT
	"=",      //ASSIGN
	"+",      //PLUS
	"-",      //MINUS
	"!",      //BANG
	"*",      //ASTERISK
	"/",      //SLASH
	"<",      //LT
	">",      //GT
	"==",     //EQ
	"!=",     //NOT_EQ
	",",      //COMMA
	";",      //SEMICOLON
	"{",      //LPAREN
	"}",      //RPAREN
	"(",      //LBRACE
	")",      //RBRACE	- works already
	"fn",     //FUNCTION
	"let",    //LET
	"true",   //TRUE
	"false",  //FALSE
	"if",     //IF
	"else",   //ELSE
	"return", //RETURN
}
