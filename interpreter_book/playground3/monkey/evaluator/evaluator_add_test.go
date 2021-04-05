package evaluator

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/token"

	"testing"
)

//TODO: specify error messages
func TestArityCallExpressions(t *testing.T) {

	tests := []struct {
		input  string
		expErr bool
		errmsg string
	}{
		// to be specified: error msg
		{"let truther = fn(x){false}; truther()", true, "not enough"},
		{"let id = fn(x){x}; id()", true, "not enough"},
		{"let add = fn(x,y){x+y}; add(2)", true, "not enough"},
		// regression
		{"let zero = fn(){0}; zero()", false, ""},
		{"let truther = fn(x){false}; truther(1)", false, ""},
		{"let id = fn(x){x}; id(1)", false, ""},
		{"let add = fn(x,y){x+y}; add(1,2)", false, ""},
		// to be specified: whether we want an error + error msg
		{"let zero = fn(){0}; zero(1)", false, ""},
		{"let zero = fn(){0}; zero(1,2,3)", false, ""},
		{"let truther = fn(x){false}; truther(1,2)", true, "too many"},
		{"let id = fn(x){x}; id(1,2)", true, "too many"},
		{"let add = fn(x,y){x+y}; add(1,2,3)", true, "too many"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		ast := p.ParseProgram()
		if len(p.Errors()) > 0 {
			t.Errorf("Unexpected parse errors for %q", tt.input) // either wrong test setup or bugs in parser
			continue
		}
		// we are testing specifically for runtime errors in evaluation
		testErrors(tt.input, ast, tt.expErr, tt.errmsg, t)
	}
}

//TODO: specify error messages
func TestDivisionByZero(t *testing.T) {

	tests := []struct {
		input  string
		expErr bool
		errmsg string
	}{
		{"3/0", true, "shame on you"},      // literally zero
		{"-3/(1-1)", true, "shame on you"}, // evaluating to zero
		{"0/1", false, ""},                 // regression
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		ast := p.ParseProgram()
		if len(p.Errors()) > 0 {
			t.Errorf("Unexpected parse errors for %q", tt.input) // either wrong test setup or bugs in parser
			continue
		}
		// we are testing specifically for runtime errors in evaluation
		testErrors(tt.input, ast, tt.expErr, tt.errmsg, t)
	}
}

func testErrors(input string, ast *ast.Program, expErr bool, expMsg string, t *testing.T) {

	env := object.NewEnvironment()

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Evaluation panics for %q: %q", input, err)
		}
	}()

	value := Eval(ast, env)
	errobj, hasErr := value.(*object.Error)

	if expErr && !hasErr {
		t.Errorf("Error missing for: %q", input)
		return
	}
	if !expErr && hasErr {
		t.Errorf("Unexpected error message for %q: %q", input, errobj.Message)
	}
	if expErr && hasErr && errobj.Message != expMsg {
		t.Errorf("Wrong error message for %q: %q (should be %q)", input, errobj.Message, expMsg)
	}
}

func TestEvalToBoolConsistency(t *testing.T) {

	env, tests := setupEvalToBoolean()

	for _, tt := range tests {
		env.Set("a", tt.object)

		expr1 := "if(a){true} else {false}"
		expr2 := "!!a"

		if evaluate(expr1, env, t) != evaluate(expr2, env, t) {
			t.Errorf("inconsistent evaluation to bool for: %q ", tt.description)
		}
	}
}

func TestEvalToBoolCorrectness(t *testing.T) {

	env, tests := setupEvalToBoolean()

	for _, tt := range tests {
		env.Set("a", tt.object)

		result := evaluate("!!a", env, t)

		switch tt.expected {
		case "true":
			if result != TRUE {
				t.Errorf("%s does not evaluate to true", tt.description)
			}
		case "false":
			if result != FALSE {
				t.Errorf("%s does not evaluate to false", tt.description)
			}
		case "error":
			if result.Type() != "ERROR" {
				t.Errorf("%s does not evaluate to an error", tt.description)
			}
		default:
			t.Errorf("test setup fails, since expectation %q not yet implemented", tt.expected)
		}
	}
}

// needs specification
func setupEvalToBoolean() (*object.Environment, []struct {
	object      object.Object
	description string
	expected    string
}) {

	// prep: create environment and build the function object (fn(){})
	env := object.NewEnvironment()
	params := []*ast.Identifier{}
	body := &ast.BlockStatement{
		Token: token.Token{Type: token.LBRACE, Literal: "{"}}
	body.Statements = []ast.Statement{}
	functionObj := &object.Function{Parameters: params, Env: env, Body: body}

	// create testdata
	tests := []struct {
		object      object.Object
		description string
		expected    string //TODO: better interface{} - but then we need to specify the errormessages!
	}{
		{TRUE, "Boolean with value true", "true"},
		{FALSE, "Boolean with value false", "false"},
		{&object.Integer{Value: -1},
			"Integer with negative value", "error"},
		{&object.Integer{Value: 0},
			"Integer with zero value", "error"},
		{&object.Integer{Value: 1},
			"Integer with positive value", "error"},
		{NULL, "Null", "error"},
		{&object.Error{Message: ""}, "Error", "error"},
		{functionObj, "Function", "error"},
		{nil, "nil", "true"},
	}
	return env, tests
}

func evaluate(input string, env *object.Environment, t *testing.T) object.Object {

	l := lexer.New(input)
	p := parser.New(l)
	ast := p.ParseProgram()
	return Eval(ast, env)
}
