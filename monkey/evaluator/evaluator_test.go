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
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
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
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
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

func TestIfElseExpression(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (0) { 10 }", nil},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		testObject(a, evaluated, tt.expected)
	}

}

func TestReturnStatements(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5;", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		testObject(a, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + true", "unknown operator: INTEGER + BOOLEAN"},
		{"5 + true; 5;", "unknown operator: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (10 > 1) { true + false; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{
			"if (10 > 1) { if (10 > 1) { return true + false; } return 1; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !a.True(ok) {
			continue
		}

		a.Equal(errObj.Message, tt.expectedMessage)
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
	case nil:
		testNilObject(a, obj)
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

func testNilObject(a *assert.Assertions, obj object.Object) {
	_, ok := obj.(*object.Null)
	a.True(ok)
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
