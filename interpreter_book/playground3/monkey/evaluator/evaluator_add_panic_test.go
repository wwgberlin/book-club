package evaluator

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"strings"
	"testing"
)

func TestPanicNotEnoughArguments(t *testing.T) {

	tests := []string{
		"let zero = fn(x){0}; zero()",
		"let id = fn(x){x}; id()",
		"let add = fn(x,y){x+y}; add(2)",
	}
	for _, input := range tests {

		l := lexer.New(input)
		p := parser.New(l)
		ast := p.ParseProgram()

		// we check specifically for a runtime error caused by the evaluation
		checkPanic(input, ast, len(p.Errors()) > 0, t)
	}
}

func TestPanicDivisionByZero(t *testing.T) {

	tests := []string{
		"3/0",
		"-3/(1-1)",
	}
	for _, input := range tests {
		l := lexer.New(input)
		p := parser.New(l)
		ast := p.ParseProgram()

		// we check specifically for a runtime error caused by the evaluation
		checkPanic(input, ast, len(p.Errors()) > 0, t)
	}
}

func TestPanicWithNil(t *testing.T) {

	nils := []string{
		"if(true){}",
		"if(false){}{let a = 1}",
		"fn(){}()",
	}
	tests := []string{
		"NIL()",
		"!NIL", // works already
		"-NIL",
		"NIL + 0",
		"0 + NIL",
		"NIL - 0",
		"0 - NIL",
		"NIL * 0",
		"0 * NIL",
		"NIL / 1",
		"0 / NIL",
		"NIL < 0",
		"0 < NIL",
		"NIL > 0",
		"0 > NIL",
		"NIL == true",
		"NIL == 0",
		"true == NIL",
		"0 == NIL",
		"NIL != true",
		"NIL != 0",
		"true != NIL",
		"0 != NIL",
	}

	for _, mynil := range nils {
		for _, schema := range tests {

			input := strings.ReplaceAll(schema, "NIL", mynil)
			l := lexer.New(input)
			p := parser.New(l)
			ast := p.ParseProgram()

			// we check specifically for a runtime error caused by the evaluation
			checkPanic(input, ast, len(p.Errors()) > 0, t)
		}
	}
}

func TestPanicWithInvalidPrograms(t *testing.T) {
	//whether that is really a problem could be discussed, since usually only valid problems are evaluated.

	// nil vs isNil !
	tests := []string{
		"let",
		"@",
		"@ let",
		"let;@; ",
	}
	for _, input := range tests {

		l := lexer.New(input)
		p := parser.New(l)
		ast := p.ParseProgram()

		// we check specifically for a runtime error caused by the evaluation
		checkPanic(input, ast, len(p.Errors()) > 0, t)
	}
}

func checkPanic(input string, ast *ast.Program, hasParseErrors bool, t *testing.T) {
	env := object.NewEnvironment()
	defer func() { // idea from https://golang.org/doc/effective_go#recover
		if err := recover(); err != nil {
			if hasParseErrors {
				t.Errorf("Panic after parse errors for %q: %q", input, err)
			} else {
				t.Errorf("Panic though no parse errors for %q: %q", input, err)
			}
		}
	}()
	Eval(ast, env)
}
