// Package ast - contains parser that builds AST for the language
package ast

import (
	"bytes"

	"github.com/technoboom/compiler/token"
)

// Node - each node in AST should implements this interface
// provides method that returns the literal value of the
// associated token
type Node interface {
	TokenLiteral() string
	String() string
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

// String - converts all the statements into string
// Creates buffer and writes all String methods exucutions
// for all child statements
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String - converts current let statement into string
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
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

// String - convers return statement into string
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

// Identifier - represents identifier
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String - returns string representation of the identifier
func (i *Identifier) String() string { return i.Value }

// ExpressionStatement - contains structure of expression
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral - returns the literal value of the associated node
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String - returns string representation of the expression
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String - returns string representation of the expression
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression - defines expression with prefix notation
// <prefix operator><expression>;
type PrefixExpression struct {
	// the prefix token, e.g. '!' or '-'
	Token token.Token
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String - returns string representation of the expression
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression - defines expression with prefix notation
// <expression><infix operator><expression>;
type InfixExpression struct {
	// the operator (infix) token, e.g. '+', '-'
	Token token.Token
	Left Expression
	Operator string
	Right Expression
}

func (oe *InfixExpression) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

// String - returns string representation of the expression
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(oe.Operator)
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}