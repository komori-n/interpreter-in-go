package token

import (
	"fmt"
)

// A kind of a token.
// We use `int` instead of `string` to enable type checks and speed up execution.
type TokenKind int

// Enumeration constants for `TokenKind`
const (
	Illegal TokenKind = iota
	Eof

	Ident
	Int

	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash

	Lt
	Gt
	Eq
	Ne

	Comma
	Semicolon

	LParen
	RParen
	LBrace
	RBrace

	Function
	Let
	True
	False
	If
	Else
	Return
)

func (tt TokenKind) String() string {
	switch tt {
	case Illegal:
		return "ILLEGAL"
	case Eof:
		return "EOF"
	case Ident:
		return "IDENT"
	case Int:
		return "INT"
	case Assign:
		return "ASSIGN"
	case Plus:
		return "PLUS"
	case Minus:
		return "MINUS"
	case Bang:
		return "BANG"
	case Asterisk:
		return "ASTERISK"
	case Slash:
		return "SLASH"
	case Lt:
		return "LT"
	case Gt:
		return "GT"
	case Eq:
		return "EQ"
	case Ne:
		return "NE"
	case Comma:
		return "COMMA"
	case Semicolon:
		return "SEMICOLON"
	case LParen:
		return "LPAREN"
	case RParen:
		return "RPAREN"
	case LBrace:
		return "LBRACE"
	case RBrace:
		return "RBRACE"
	case Function:
		return "FUNCTION"
	case Let:
		return "LET"
	case True:
		return "TRUE"
	case False:
		return "FALSE"
	case If:
		return "IF"
	case Else:
		return "ELSE"
	case Return:
		return "RETURN"
	default:
		return fmt.Sprintf("%d", int(tt))
	}
}

var keywards = map[string]TokenKind{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookUpIdent(ident string) TokenKind {
	if tok, ok := keywards[ident]; ok {
		return tok
	}
	return Ident
}

// A code token
type Token struct {
	Kind    TokenKind
	Literal string
	Line    int
}
