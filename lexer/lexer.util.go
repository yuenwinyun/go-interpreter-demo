package lexer

import "go-interpreter-demo/token"

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func createToken(Type token.Type, ch byte) token.Token {
	return token.Token{Type: Type, Literal: string(ch)}
}
