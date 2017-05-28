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

// parseStatement - parses the statement to make a decision what kind of
// statement it can be, after calls the appropriate function to perform
// correct action with the statement and return ast.Statement object
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement - parses
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs - checks if current token type is a given type
func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

// peekTokenIs - checks if peek token type is a given type
func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

// expectPeek - peeks the token if types match, otherwise, returns false
func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

// ParseProgram - parses root node and produces AST
// Build the root node of the AST
// After this it reads tokens one by one until he reached token.EOF
// On each iteration it does parsing statement, if it's success - this statement
// adds to the Statements of the program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	// read until we reaced the end of the file
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// proceed with next token
		p.nextToken()
	}
	return program
}
