// Package parser - contains Parser which transforms tokens into AST
package parser

import (
	"github.com/technoboom/compiler/ast"
	"github.com/technoboom/compiler/lexer"
	"github.com/technoboom/compiler/token"
)

// Parser - structure for storing lexer and state of parsing
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New - creates new Parser accordingly to the lexer in the args
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens to ensure that curToken and peekToken are
	// both set
	p.nextToken()
	p.nextToken()

	return p
}

// Reads next token from the Lexer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram - parses root node and produces AST
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
