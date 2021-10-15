package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {
	a := assert.New(t)

	input := `
let x = 5;
let y = 10;
let foobar = 334334;`
	tests := []struct {
		expectedIdentifier string
		expectedValue      int
	}{
		{"x", 5},
		{"y", 10},
		{"foobar", 334334},
	}

	program := parse(a, input)
	if !a.Equal(len(program.Statements), len(tests)) {
		return
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		a.Equal(stmt.TokenLiteral(), "let")
		letStmt, ok := stmt.(*ast.LetStatement)
		if !a.True(ok, "*ast.LetStatement") {
			continue
		}
		a.Equal(letStmt.Name.Value, tt.expectedIdentifier)
		a.Equal(letStmt.Name.TokenLiteral(), tt.expectedIdentifier)
		testLiteralExpression(a, letStmt.Value, tt.expectedValue)
	}
}

func TestReturnStatements(t *testing.T) {
	a := assert.New(t)

	input := `
return 5;
return 10;
return 993322;`
	tests := []struct {
		expectedValue int
	}{
		{5}, {10}, {993322},
	}
	program := parse(a, input)
	if !a.Equal(len(program.Statements), len(tests)) {
		return
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		a.Equal(stmt.TokenLiteral(), "return")
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !a.True(ok, "*ast.ReturnStatement") {
			continue
		}
		a.Equal(returnStmt.TokenLiteral(), "return")
		testLiteralExpression(a, returnStmt.ReturnValue, tt.expectedValue)
	}
}

func TestParsingArrayLiterals(t *testing.T) {
	a := assert.New(t)

	input := "[1, 3, 9]"
	program := parse(a, input)
	if !a.Equal(len(program.Statements), 1) {
		return
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok) {
		return
	}
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !a.True(ok) {
		return
	}

	a.Equal(len(array.Elements), 3)
	testLiteralExpression(a, array.Elements[0], 1)
	testLiteralExpression(a, array.Elements[1], 3)
	testLiteralExpression(a, array.Elements[2], 9)
}

func TestParsingHashLiteral(t *testing.T) {
	a := assert.New(t)
	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`
	program := parse(a, input)
	if !a.Equal(len(program.Statements), 1) {
		return
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok) {
		return
	}
	_, ok = stmt.Expression.(*ast.HashLiteral)
	if !a.True(ok) {
		return
	}
	// a.Equal(hash.Pairs[&ast.StringLiteral{Value: "one"}].String(), "0 + 1")
}

func TestParsingIndexExpressions(t *testing.T) {
	a := assert.New(t)
	input := "myArray[2]"
	program := parse(a, input)
	if !a.Equal(len(program.Statements), 1) {
		return
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok) {
		return
	}
	ie, ok := stmt.Expression.(*ast.IndexExpression)
	if !a.True(ok) {
		return
	}
	testIdendifier(a, ie.Left, "myArray")
	testLiteralExpression(a, ie.Index, 2)
}

func TestStringLiteralExpression(t *testing.T) {
	a := assert.New(t)
	input := `"Hello World";`
	program := parse(a, input)

	if !a.Equal(len(program.Statements), 1) {
		return
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok, "*ast.ExpressionStatement") {
		return
	}
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !a.True(ok, "*ast.StringLiteral") {
		return
	}
	a.Equal(literal.Value, "Hello World")
}

func TestIntegerLiteralExpression(t *testing.T) {
	a := assert.New(t)
	input := "334;"
	program := parse(a, input)
	if !a.Equal(len(program.Statements), 1) {
		return
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok, "*ast.ExpressionStatement") {
		return
	}
	testLiteralExpression(a, stmt.Expression, 334)
}

func TestPrefixExpressions(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input        string
		operator     string
		integerValue int
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range tests {
		program := parse(a, tt.input)
		if !a.Equal(len(program.Statements), 1) {
			continue
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !a.True(ok, "*ast.ExpressionStatement") {
			continue
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !a.True(ok, "*ast.PrefixExpression") {
			continue
		}
		a.Equal(exp.Operator, tt.operator)
		testLiteralExpression(a, exp.Right, tt.integerValue)
	}
}

func TestInfixExpressions(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input      string
		leftValue  int
		operator   string
		rightValue int
	}{
		{"33 + 4", 33, "+", 4},
		{"33 - 4", 33, "-", 4},
		{"33 * 4", 33, "*", 4},
		{"33 / 4", 33, "/", 4},
		{"33 > 4", 33, ">", 4},
		{"33 < 4", 33, "<", 4},
		{"33 == 4", 33, "==", 4},
		{"33 != 4", 33, "!=", 4},
	}

	for _, tt := range tests {
		program := parse(a, tt.input)
		if !a.Equal(len(program.Statements), 1) {
			continue
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !a.True(ok, "*ast.ExpressionStatement") {
			continue
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !a.True(ok, "*ast.InfixExpression") {
			continue
		}
		testLiteralExpression(a, exp.Left, tt.leftValue)
		a.Equal(exp.Operator, tt.operator)
		testLiteralExpression(a, exp.Right, tt.rightValue)
	}
}

func TestBooleanExpression(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input         string
		expectedValue bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		program := parse(a, tt.input)
		if !a.Equal(len(program.Statements), 1) {
			continue
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !a.True(ok, "*ast.ExpressionStatement") {
			continue
		}
		exp, ok := stmt.Expression.(*ast.Boolean)
		if !a.True(ok, "*ast.Boolean") {
			continue
		}
		testLiteralExpression(a, exp, tt.expectedValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"(5 + 5) * 2 * (5 + 5)", "(((5 + 5) * 2) * (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
	}

	for _, tt := range tests {
		program := parse(a, tt.input)
		a.Equal(program.String(), tt.expected)
	}
}

func TestIfExpression(t *testing.T) {
	a := assert.New(t)
	input := `if (x < y) { x }`
	program := parse(a, input)

	if a.Equal(len(program.Statements), 1) {
		return
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok, "*ast.ExpressionStatement") {
		return
	}
	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !a.True(ok, "*ast.IfExpression") {
		return
	}
	a.Equal(exp.Condition.String(), "x < y")
	a.Equal(exp.Consequence.String(), "x")
	a.Nil(exp.Alternative)
}

func TestElseExpression(t *testing.T) {
	a := assert.New(t)
	input := `if (x < y) { x } else { y }`
	program := parse(a, input)

	if a.Equal(len(program.Statements), 1) {
		return
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok, "*ast.ExpressionStatement") {
		return
	}
	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !a.True(ok, "*ast.IfExpression") {
		return
	}
	a.Equal(exp.Condition.String(), "x < y")
	a.Equal(exp.Consequence.String(), "x")
	a.Equal(exp.Alternative.String(), "y")
}

func TestFunctionLiteralParsing(t *testing.T) {
	a := assert.New(t)
	input := `fn(x, y) { x + y; }`
	program := parse(a, input)

	if a.Equal(len(program.Statements), 1) {
		return
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !a.True(ok, "*ast.ExpressionStatement") {
		return
	}
	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !a.True(ok, "*ast.FunctionLiteral") {
		return
	}
	if !a.Equal(len(function.Parameters), 2) {
		return
	}
	testLiteralExpression(a, function.Parameters[0], "x")
	testLiteralExpression(a, function.Parameters[1], "y")

	if !a.Equal(len(function.Body.Statements), 1) {
		return
	}
	a.Equal(function.Body.Statements[0].String(), "x + y")
}

func parse(a *assert.Assertions, input string) *ast.Program {
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(a, p)
	a.NotNil(program)

	return program
}

func checkParserErrors(a *assert.Assertions, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	for _, msg := range errors {
		a.Failf("parser error", msg)
	}
}

func testLiteralExpression(a *assert.Assertions, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteralExpression(a, exp, int64(v))
	case int64:
		return testIntegerLiteralExpression(a, exp, v)
	case string:
		return testIdendifier(a, exp, v)
	case bool:
		return testBooleanLiteral(a, exp, v)
	}
	a.Fail("type of exp not handled")
	return false
}

func testIntegerLiteralExpression(a *assert.Assertions, exp ast.Expression, value int64) bool {
	integ, ok := exp.(*ast.IntegerLiteral)
	if !a.True(ok, "*ast.IntegerLiteral") {
		return false
	}

	if !a.Equal(integ.Value, value) {
		return false
	}
	if !a.Equal(integ.TokenLiteral(), fmt.Sprintf("%d", value)) {
		return false
	}
	return true
}

func testIdendifier(a *assert.Assertions, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !a.True(ok, "*ast.Identifier") {
		return false
	}
	if !a.Equal(ident.Value, value) {
		return false
	}
	if !a.Equal(ident.TokenLiteral(), value) {
		return false
	}
	return true
}

func testBooleanLiteral(a *assert.Assertions, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !a.True(ok, "*ast.Identifier") {
		return false
	}
	if !a.Equal(bo.Value, value) {
		return false
	}
	if !a.Equal(bo.TokenLiteral(), fmt.Sprintf("%t", value)) {
		return false
	}
	return true
}
