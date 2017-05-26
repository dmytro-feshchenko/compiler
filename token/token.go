// Package token - Contains all tokens for the lexical analysis
package token

const (
	// ILLEGAL - identifies a token/character we don't know
	ILLEGAL = "ILLEGAL"
	// EOF - end of file
	EOF = "EOF"

	// IDENT - identifier
	IDENT = "IDENT"
	// INT - integer number
	INT = "INT"

	// LPAREN - left parenthesis
	LPAREN = "("
	// RPAREN - right parenthesis
	RPAREN = ")"
	// LBRACKET - left (open) bracket
	LBRACKET = "{"
	// RBRACKET - right (close) bracket
	RBRACKET = "}"

	// COMMA - comma between operands, declarations, etc.
	COMMA = ","
	// SEMICOLON - semicolon between expressions
	SEMICOLON = ";"

	// PLUS - add/concat operator
	PLUS = "+"
	// MINUS - minus/negative operator
	MINUS = "-"
	// MULTIPLY - multiply operator
	MULTIPLY = "*"
	// DIVIDE - division operator
	DIVIDE = "/"
	// ASSIGN - assiang operator
	ASSIGN = "="

	// LET - let keyword
	LET = "LET"
	// FUNCTION - function keyword
	FUNCTION = "FUNCTION"
)

// Type - contains token type
type Type string

// Token - contains the type and literal of the language token
type Token struct {
	Type    Type
	Literal string
}
