package parser

import (
	"fmt"
	"go-interpreter-demo/token"
)

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) appendError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.nextToken.Type)
	p.errors = append(p.errors, msg)
}
