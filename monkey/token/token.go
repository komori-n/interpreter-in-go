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
		return "="
	case Plus:
		return "+"
	case Minus:
		return "-"
	case Bang:
		return "!"
	case Asterisk:
		return "*"
	case Slash:
		return "/"
	case Lt:
		return "<"
	case Gt:
		return ">"
	case Eq:
		return "=="
	case Ne:
		return "!="
	case Comma:
		return ","
	case Semicolon:
		return ";"
	case LParen:
		return "("
	case RParen:
		return ")"
	case LBrace:
		return "{"
	case RBrace:
		return "}"
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

// Judge if the argument is a keyword or not.
// If so, return the kind of keyword. Else, return `Ident`.
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
