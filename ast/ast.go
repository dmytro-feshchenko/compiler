// Package ast - contains parser that builds AST for the language
package ast

import "github.com/technoboom/compiler/token"

// Node - each node in AST should implements this interface
// provides method that returns the literal value of the
// associated token
type Node interface {
	TokenLiteral() string
}

// Statement - subset of nodes which represents statements
type Statement interface {
	Node
	statementNode()
}

// Expression - subset of nodes which represents expressions
type Expression interface {
	Node
	expressionNode()
}

// Program - root node for program AST
type Program struct {
	Statements []Statement
}

// TokenLiteral - returns the literal value of the associated node
// if program contains more than 1 statement - returns the first one
// else returns empty string
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement - statement that represents sentences with let
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral - returns the literal value of the associated node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement - statement that represents return <expression>;
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral - returns the literal value of the associated node
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// Identifier - represents identifier
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
