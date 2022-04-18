package lexer

import (
	"go-interpreter-demo/core/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input string
	var testCases []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}
	var lexer *Lexer

	input = "=+(){},;"
	testCases = []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.RBRACKET, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	lexer = New(input)
	for index, test := range testCases {
		token := lexer.NextToken()
		if token.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", index, test.ExpectedType, token.Type)
		}
		if token.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", index, test.ExpectedLiteral, token.Literal)
		}
	}

	input = `
		let five = 5; 
		let ten = 10;
		let add = fn(x, y) {
			x + y; 
		};
		let result = add(five, ten);
	`
	testCases = []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lexer = New(input)
	for index, test := range testCases {
		tok := lexer.NextToken()

		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", index, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", index, test.ExpectedLiteral, tok.Literal)
		}
	}

	input = `
		!-/*5; 
		5 < 10 > 5;
	`
	testCases = []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lexer = New(input)
	for index, test := range testCases {
		tok := lexer.NextToken()

		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", index, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", index, test.ExpectedLiteral, tok.Literal)
		}
	}

	input = `
		if (5 < 10) {
			return true; 
		} else {
			return false; 
		}
	`
	testCases = []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.ELSE, "else"},
		{token.LBRACKET, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
	}
	lexer = New(input)
	for index, test := range testCases {
		token := lexer.NextToken()
		if token.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", index, test.ExpectedType, token.Type)
		}
		if token.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", index, test.ExpectedLiteral, token.Literal)
		}
	}

	input = `
		10 == 10;
		10 != 9;
	`

	testCases = []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}
	lexer = New(input)
	for index, test := range testCases {
		token := lexer.NextToken()
		if token.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", index, test.ExpectedType, token.Type)
		}
		if token.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", index, test.ExpectedLiteral, token.Literal)
		}
	}
}
