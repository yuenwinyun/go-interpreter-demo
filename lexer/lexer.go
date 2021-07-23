package lexer

import "go-interpreter-demo/token"

type Lexer struct {
	input           string
	ch              byte
	currentPosition int
	nextPosition    int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.moveChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		if l.nextChar() == '=' {
			ch := l.ch
			l.moveChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = createToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = createToken(token.PLUS, l.ch)
	case '-':
		tok = createToken(token.MINUS, l.ch)
	case '!':
		if l.nextChar() == '=' {
			ch := l.ch
			l.moveChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = createToken(token.NOT, l.ch)
		}
	case '/':
		tok = createToken(token.SLASH, l.ch)
	case '*':
		tok = createToken(token.ASTERISK, l.ch)
	case '>':
		tok = createToken(token.GT, l.ch)
	case '<':
		tok = createToken(token.LT, l.ch)
	case ';':
		tok = createToken(token.SEMICOLON, l.ch)
	case ',':
		tok = createToken(token.COMMA, l.ch)
	case '(':
		tok = createToken(token.LPAREN, l.ch)
	case ')':
		tok = createToken(token.RPAREN, l.ch)
	case '{':
		tok = createToken(token.LBRACE, l.ch)
	case '}':
		tok = createToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = createToken(token.ILLEGAL, l.ch)
		}
	}

	l.moveChar()

	return tok
}

func (l *Lexer) nextChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) moveChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currentPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.ch) {
		l.moveChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.ch) {
		l.moveChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.moveChar()
	}
}
