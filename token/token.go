// Package token - Contains all tokens for the lexical analysis
package token

var keywords = map[string]Type{
	"function": FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

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
	// ASSIGN - assign operator
	ASSIGN = "="
	// BANG - logical not, symbol for inversion
	BANG = "!"
	// LT - less than smth. (logical operator)
	LT = "<"
	// GT - great than smth. (logical operator)
	GT = ">"
	// OR - OR operator
	OR = "|"
	// EQ - operator for checking if both operands are equal
	EQ = "=="
	// NOTEQ - operator for checking if both operands are not equal
	NOTEQ = "!="

	// LET - let keyword
	LET = "LET"
	// FUNCTION - function keyword
	FUNCTION = "FUNCTION"
	// TRUE - boolean true
	TRUE = "TRUE"
	// FALSE - boolean false
	FALSE = "FALSE"
	// IF - condition
	IF = "IF"
	// ELSE - else condition
	ELSE = "ELSE"
	// RETURN - return keyword for the function
	RETURN = "RETURN"
)

// Type - contains token type
type Type string

// Token - contains the type and literal of the language token
type Token struct {
	Type    Type
	Literal string
}

// LookupIdent - looks into the keywords map to check if
// the given identifier is a keyword
// If it's a keyword - returns the keyword token type
// Otherwise, returns token.IDENT which is used for all user
// defined identifiers
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
