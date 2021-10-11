package evaluator

import (
	"testing"

	"monkey/lexer"
	"monkey/object"
	"monkey/parser"

	"github.com/stretchr/testify/assert"
)

func TestEvalIntegerExpression(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		testObject(a, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		testObject(a, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!0", true},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		testObject(a, evaluated, tt.expected)
	}
}

func testObject(a *assert.Assertions, obj object.Object, expected interface{}) {
	switch v := expected.(type) {
	case int:
		testIntegerObject(a, obj, int64(v))
	case int64:
		testIntegerObject(a, obj, v)
	case bool:
		testBooleanObject(a, obj, v)
	default:
		a.Fail("type of obj not handles")
	}
}

func testIntegerObject(a *assert.Assertions, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	if !a.True(ok) {
		return
	}
	a.Equal(result.Value, expected)
}

func testBooleanObject(a *assert.Assertions, obj object.Object, expected bool) {
	result, ok := obj.(*object.Boolean)
	if !a.True(ok) {
		return
	}
	a.Equal(result.Value, expected)
}

func testEval(a *assert.Assertions, input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	checkParserErrors(a, p)
	a.NotNil(program)

	return Eval(program)
}

func checkParserErrors(a *assert.Assertions, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	for _, msg := range errors {
		a.Failf("parser error", msg)
	}
}
