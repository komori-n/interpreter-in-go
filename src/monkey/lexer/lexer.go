package lexer

import (
	"monkey/token"
	"unicode"
)

type Lexer struct {
	input        []rune
	position     int
	line         int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input), line: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	if l.ch == '\n' {
		l.line += 1
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.Assign, l.ch, l.line)
	case ';':
		tok = newToken(token.Semicolon, l.ch, l.line)
	case '(':
		tok = newToken(token.LParen, l.ch, l.line)
	case ')':
		tok = newToken(token.RParen, l.ch, l.line)
	case ',':
		tok = newToken(token.Comma, l.ch, l.line)
	case '+':
		tok = newToken(token.Plus, l.ch, l.line)
	case '{':
		tok = newToken(token.LBrace, l.ch, l.line)
	case '}':
		tok = newToken(token.RBrace, l.ch, l.line)
	case 0:
		// Assign empty string instead of null string ("\0")
		tok.Kind = token.Eof
		tok.Literal = ""
		tok.Line = l.line
	default:
		tok.Line = l.line
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Kind = token.LookUpIdent(tok.Literal)
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Kind = token.Int
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.Illegal, l.ch, l.line)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenKind token.TokenKind, ch rune, line int) token.Token {
	return token.Token{Kind: tokenKind, Literal: string(ch), Line: line}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
