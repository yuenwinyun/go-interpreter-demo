package lexer

import (
	utils "go-interpreter-demo/core/common"
	"go-interpreter-demo/core/token"
)

type Lexer struct {
	ch           byte
	input        string
	position     int
	readPosition int
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = createToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = createToken(token.BANG, l.ch)
		}
	case '+':
		tok = createToken(token.PLUS, l.ch)
	case '-':
		tok = createToken(token.MINUS, l.ch)
	case '>':
		tok = createToken(token.GT, l.ch)
	case '<':
		tok = createToken(token.LT, l.ch)
	case '*':
		tok = createToken(token.ASTERISK, l.ch)
	case '/':
		tok = createToken(token.SLASH, l.ch)
	case '(':
		tok = createToken(token.LPAREN, l.ch)
	case ')':
		tok = createToken(token.RPAREN, l.ch)
	case '{':
		tok = createToken(token.LBRACKET, l.ch)
	case '}':
		tok = createToken(token.RBRACKET, l.ch)
	case ',':
		tok = createToken(token.COMMA, l.ch)
	case ';':
		tok = createToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if utils.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if utils.IsDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = createToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for utils.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func createToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}
