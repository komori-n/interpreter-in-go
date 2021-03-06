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

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = newToken(token.Eq, literal, l.line)
		} else {
			tok = newToken(token.Assign, string(l.ch), l.line)
		}
	case ':':
		tok = newToken(token.Colon, string(l.ch), l.line)
	case ';':
		tok = newToken(token.Semicolon, string(l.ch), l.line)
	case '(':
		tok = newToken(token.LParen, string(l.ch), l.line)
	case ')':
		tok = newToken(token.RParen, string(l.ch), l.line)
	case ',':
		tok = newToken(token.Comma, string(l.ch), l.line)
	case '+':
		tok = newToken(token.Plus, string(l.ch), l.line)
	case '-':
		tok = newToken(token.Minus, string(l.ch), l.line)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = newToken(token.Ne, literal, l.line)
		} else {
			tok = newToken(token.Bang, string(l.ch), l.line)
		}
	case '*':
		tok = newToken(token.Asterisk, string(l.ch), l.line)
	case '/':
		tok = newToken(token.Slash, string(l.ch), l.line)
	case '<':
		tok = newToken(token.Lt, string(l.ch), l.line)
	case '>':
		tok = newToken(token.Gt, string(l.ch), l.line)
	case '{':
		tok = newToken(token.LBrace, string(l.ch), l.line)
	case '}':
		tok = newToken(token.RBrace, string(l.ch), l.line)
	case '[':
		tok = newToken(token.LBracket, string(l.ch), l.line)
	case ']':
		tok = newToken(token.RBracket, string(l.ch), l.line)
	case '"':
		tok.Line = l.line
		tok.Kind = token.String
		tok.Literal = l.readString()
	case 0:
		// Assign empty string instead of null string ("\0")
		tok = newToken(token.Eof, "", l.line)
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
			tok = newToken(token.Illegal, string(l.ch), l.line)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenKind token.TokenKind, literal string, line int) token.Token {
	return token.Token{Kind: tokenKind, Literal: literal, Line: line}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return string(l.input[position:l.position])
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
