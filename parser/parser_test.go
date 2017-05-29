package parser

import (
	"testing"

	"github.com/technoboom/compiler/ast"
	"github.com/technoboom/compiler/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
  let x = 100;
  let y = 10;
  let result = true;
  `
	// create new lexer for the input
	l := lexer.New(input)
	// create parser with the lexer
	p := New(l)
	checkParserErrors(t, p)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parser.ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements len mismatch, expected=%q, got=%q",
			3, len(program.Statements),
		)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"result"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() is not 'let', got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%q", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got=%q", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name not '%s', got=%q", name, letStmt.Name)
		return false
	}

	return true
}

// checkParserErrors - checks the erros of the parser
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("The parser has %d errors, must be 0", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func TestParserErrors(t *testing.T) {
	input := `
	let 8999;
	let 1 = a;
	let = 5;
	`

	l := lexer.New(input)
	p := New(l)
	p.ParseProgram()
	errors := p.Errors()
	if len(errors) != 3 {
		t.Errorf("Parser has %d errors, expected 3", len(errors))
	}
	// tests := []struct {
	// 	expectedError string
	// }{
	// 	{"expected next token to be IDENT, got INT instead"},
	// 	{"expected next token to be IDENT, got INT instead"},
	// 	{"expected next token to be IDENT, got ASSIGN instead"},
	// }
	//
	// for i, tt := range tests {
	// 	if tt.expectedError != errors[i] {
	// 		t.Errorf("", args)
	// 	}
	// }
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return a;
	return 5123;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements len mismatch, expected=%d, got=%d",
			3, len(program.Statements),
		)
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("stmt TokenLiteral is not 'return', got=%q",
				returnStmt.TokenLiteral())
		}
	}
}
