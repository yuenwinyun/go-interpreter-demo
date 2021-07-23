package parser

import (
	"fmt"
	"go-interpreter-demo/ast"
	"go-interpreter-demo/lexer"
	"go-interpreter-demo/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESS_OR_GREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.Type]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESS_OR_GREATER,
	token.GT:       LESS_OR_GREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

type (
	prefixParserFunc func() ast.Expression
	infixParserFunc  func(ast.Expression) ast.Expression
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
	errors       []string

	prefixParserFns map[token.Type]prefixParserFunc
	infixParserFns  map[token.Type]infixParserFunc
}

func New(sourceCode string) *Parser {
	l := lexer.New(sourceCode)
	p := &Parser{l: l, errors: []string{}}

	p.prefixParserFns = make(map[token.Type]prefixParserFunc)
	p.infixParserFns = make(map[token.Type]infixParserFunc)
	p.registerPrefixParserFunc(token.IDENT, p.parseIdentifier)
	p.registerPrefixParserFunc(token.INT, p.parseInteger)
	p.registerPrefixParserFunc(token.NOT, p.parsePrefixExpression)
	p.registerPrefixParserFunc(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixParserFunc(token.TRUE, p.parseBoolean)
	p.registerPrefixParserFunc(token.FALSE, p.parseBoolean)
	p.registerPrefixParserFunc(token.LPAREN, p.parseGroupedExpression)

	for _, tok := range []token.Type{token.MINUS, token.SLASH, token.ASTERISK, token.EQ, token.NOT_EQ, token.LT, token.GT} {
		p.registerInfixParserFunc(tok, p.parseInfixExpression)
	}

	p.getTokenReady()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.moveToken()
	}

	return program
}

func (p *Parser) getTokenReady() {
	p.moveToken()
	p.moveToken()
}

func (p *Parser) registerPrefixParserFunc(tokenType token.Type, fn prefixParserFunc) {
	p.prefixParserFns[tokenType] = fn
}

func (p *Parser) registerInfixParserFunc(tokenType token.Type, fn infixParserFunc) {
	p.infixParserFns[tokenType] = fn
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.currentToken, Value: p.currentTokenIs(token.TRUE)}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.moveToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) moveToken() {
	p.currentToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) currentTokenIs(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p *Parser) nextTokenIs(t token.Type) bool {
	return p.nextToken.Type == t
}

func (p *Parser) expectNextTokenIs(t token.Type) bool {
	if p.nextTokenIs(t) {
		p.moveToken()
		return true
	} else {
		p.appendError(t)
		return false
	}
}

func (p *Parser) parseInteger() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("count not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.moveToken()

	if expression.Operator == "+" {
		expression.Right = p.parseExpression(precedence - 1)
	} else {
		expression.Right = p.parseExpression(precedence)
	}

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.moveToken()

	expression := p.parseExpression(LOWEST)

	if !p.expectNextTokenIs(token.RPAREN) {
		return nil
	}

	return expression
}

func (p *Parser) nextPrecedence() int {
	if p, ok := precedences[p.nextToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}