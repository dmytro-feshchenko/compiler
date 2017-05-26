package lexer

import "github.com/technoboom/compiler/token"

// Lexer - contains data and methods for performing lexical analysis
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	character    byte // current char after examination
}

// New - creates new lexer with give input
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// readChar - reads current character from readPosition into character var
// shifts readPosition at 1 position forward
// if the readPosition out of input len, set character to 0
func (l *Lexer) readChar() {
	// check if we not out of input len
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}
	l.position = l.readPosition
	// shift caret at 1 position forward
	l.readPosition++
}

// NextToken - parse next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.character {
	case '=':
		tok = newToken(token.ASSIGN, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case '{':
		tok = newToken(token.LBRACKET, l.character)
	case '}':
		tok = newToken(token.RBRACKET, l.character)
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '-':
		tok = newToken(token.MINUS, l.character)
	case '*':
		tok = newToken(token.MULTIPLY, l.character)
	case '/':
		tok = newToken(token.DIVIDE, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// newToken - creates new token with given type and literal
func newToken(tokenType token.Type, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}
