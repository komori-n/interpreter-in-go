package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
3 < 3 > 4;

if (5 < 10) {
	return true;
} else {
	return false;
}
`

	tests := []struct {
		expectedKind    token.TokenKind
		expectedLiteral string
		expectedLine    int
	}{
		// 1st line
		{token.Let, "let", 1},
		{token.Ident, "five", 1},
		{token.Assign, "=", 1},
		{token.Int, "5", 1},
		{token.Semicolon, ";", 1},
		// 2nd line
		{token.Let, "let", 2},
		{token.Ident, "ten", 2},
		{token.Assign, "=", 2},
		{token.Int, "10", 2},
		{token.Semicolon, ";", 2},

		// 3rd -- 5th line
		{token.Let, "let", 4},
		{token.Ident, "add", 4},
		{token.Assign, "=", 4},
		{token.Function, "fn", 4},
		{token.LParen, "(", 4},
		{token.Ident, "x", 4},
		{token.Comma, ",", 4},
		{token.Ident, "y", 4},
		{token.RParen, ")", 4},
		{token.LBrace, "{", 4},
		{token.Ident, "x", 5},
		{token.Plus, "+", 5},
		{token.Ident, "y", 5},
		{token.Semicolon, ";", 5},
		{token.RBrace, "}", 6},
		{token.Semicolon, ";", 6},

		// 6th line
		{token.Let, "let", 8},
		{token.Ident, "result", 8},
		{token.Assign, "=", 8},
		{token.Ident, "add", 8},
		{token.LParen, "(", 8},
		{token.Ident, "five", 8},
		{token.Comma, ",", 8},
		{token.Ident, "ten", 8},
		{token.RParen, ")", 8},
		{token.Semicolon, ";", 8},

		// 7th line
		{token.Bang, "!", 9},
		{token.Minus, "-", 9},
		{token.Slash, "/", 9},
		{token.Asterisk, "*", 9},
		{token.Int, "5", 9},
		{token.Semicolon, ";", 9},

		// 8th line
		{token.Int, "3", 10},
		{token.Lt, "<", 10},
		{token.Int, "3", 10},
		{token.Gt, ">", 10},
		{token.Int, "4", 10},
		{token.Semicolon, ";", 10},

		// 9th --  line
		{token.If, "if", 12},
		{token.LParen, "(", 12},
		{token.Int, "5", 12},
		{token.Lt, "<", 12},
		{token.Int, "10", 12},
		{token.RParen, ")", 12},
		{token.LBrace, "{", 12},
		{token.Return, "return", 13},
		{token.True, "true", 13},
		{token.Semicolon, ";", 13},
		{token.RBrace, "}", 14},
		{token.Else, "else", 14},
		{token.LBrace, "{", 14},
		{token.Return, "return", 15},
		{token.False, "false", 15},
		{token.Semicolon, ";", 15},
		{token.RBrace, "}", 16},

		{token.Eof, "", 17},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Kind != tt.expectedKind {
			t.Fatalf("tests[%d] - tokenkind wrong(L%d). expected=%q, got=%q", i, tt.expectedLine, tt.expectedKind, tok.Kind)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong(L%d). expected=%q, got=%q", i, tt.expectedLine, tt.expectedLiteral, tok.Literal)
		}

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong(L%d). expected=%d, got=%d", i, tt.expectedLine, tt.expectedLine, tok.Line)
		}
	}
}
