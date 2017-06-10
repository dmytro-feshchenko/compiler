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

	l.skipWhitespace()

	switch l.character {
	case '=':
		// look ahead on 1 position to check if it's not the ==
		if l.pickChar() == '=' {
			character := l.character
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(character) + string(l.character),
			}
		} else {
			tok = newToken(token.ASSIGN, l.character)
		}
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
	case '<':
		tok = newToken(token.LT, l.character)
	case '>':
		tok = newToken(token.GT, l.character)
	case '!':
		// look ahead on 1 position to check if it's not the !=
		if l.pickChar() == '=' {
			character := l.character
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(character) + string(l.character),
			}
		} else {
			tok = newToken(token.BANG, l.character)
		}
	case ',':
		tok = newToken(token.COMMA, l.character)
	case '|':
		tok = newToken(token.OR, l.character)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.character) {
			// read keyword or identifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.character) {
			// read number
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		// illegal token
		tok = newToken(token.ILLEGAL, l.character)
	}
	l.readChar()
	return tok
}

// newToken - creates new token with given type and literal
func newToken(tokenType token.Type, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.character) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Ignore all spaces, tabulation and line breakers
func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' ||
		l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

// reads the whole number while the next symbol is digit
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.character) {
		l.readChar()
	}
	return l.input[position:l.position]
}


// readString - reads string until double quote will not be met
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.character == '"' {
			break
		}
	}
	return l.input[position:l.position]
}

// pickChar - picks one character if the position of carriage
// is not out of input len
func (l *Lexer) pickChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// Checks if the character is English letter or underscore
func isLetter(character byte) bool {
	return ('a' <= character && character <= 'z') ||
		('A' <= character && character <= 'Z') ||
		character == '_'
}

// Checks if the character is digit (0-9)
func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}