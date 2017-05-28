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
	parser := New(l)

	program := parser.ParseProgram()
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
