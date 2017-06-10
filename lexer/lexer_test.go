package lexer

import (
	"testing"

	"github.com/technoboom/compiler/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = function(x, y) {
		x + y;
	}

	let result = add(five, ten);
	let percent = (five - ten) / result * 100;
	let boolVar = !(five < ten || result < ten);
	if (5 < 10) {
		return false;
	} else {
		return true;
	}

	if (true == false) {
		return false;
	} else if (true != 1) {
		return true;
	}
	"hello"
	"hello, world!"
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "percent"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.MINUS, "-"},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.DIVIDE, "/"},
		{token.IDENT, "result"},
		{token.MULTIPLY, "*"},
		{token.INT, "100"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "boolVar"},
		{token.ASSIGN, "="},
		{token.BANG, "!"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.LT, "<"},
		{token.IDENT, "ten"},
		{token.OR, "|"},
		{token.OR, "|"},
		{token.IDENT, "result"},
		{token.LT, "<"},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.ELSE, "else"},
		{token.LBRACKET, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.EQ, "=="},
		{token.FALSE, "false"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.EQ, "!="},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.STRING, "hello"},
		{token.STRING, "hello, world!"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

	// check EOF
	eofToken := l.NextToken()
	if eofToken.Type != token.EOF {
		t.Fatalf("test for EOF failed. expected=EOF, got=%q", eofToken.Type)
	}
}
