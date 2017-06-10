// Package ast - contains parser that builds AST for the language
package ast

import (
	"bytes"

	"github.com/technoboom/compiler/token"
	"strings"
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
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean - structure for boolean variables
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String - returns string representation of the expression
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression - defines conditional expression
// if (<condition>) <consequence> else <alternative>
type IfExpression struct {
	Token token.Token // the `if` token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String - returns string representation of the expression
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement - represents series of statements
type BlockStatement struct {
	// the "{" token
	Token token.Token
	// the series of statements
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral - returns the literal value of the associated node
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String - returns string representation of the expression
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral - represents functions
type FunctionLiteral struct {
	// The 'function' token
	Token token.Token

	Parameters []*Identifier
	Body *BlockStatement
}


func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String - returns string representation of the expression
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	// collect all parameters of function as strings
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression - defines calls of expressions
// <expression>(<comma separated expressions>)
type CallExpression struct {
	// the `(` token
	Token token.Token

	Function Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String - returns string representation of the expression
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// StringLiteral - represents strings
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral - returns the literal value of the associated node
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// String - returns string representation of the expression
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}