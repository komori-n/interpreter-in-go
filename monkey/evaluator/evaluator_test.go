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
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(a, tt.input)
		testIntegerObject(a, evaluated, tt.expected)
	}
}

func testIntegerObject(a *assert.Assertions, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
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
