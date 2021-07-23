package parser

import (
	"fmt"
	"go-interpreter-demo/ast"
	"go-interpreter-demo/token"
)

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParserFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.nextTokenIs(token.SEMICOLON) && precedence < p.nextPrecedence() {
		infix := p.infixParserFns[p.nextToken.Type]
		if infix == nil {
			return leftExp
		}

		p.moveToken()

		leftExp = infix(leftExp)
	}
	return leftExp
}
