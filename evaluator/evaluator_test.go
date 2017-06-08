package evaluator

import (
	"testing"
	"github.com/technoboom/compiler/object"
	"github.com/technoboom/compiler/lexer"
	"github.com/technoboom/compiler/parser"
)

func TestEvalIntegerExpressions(t *testing.T) {
	tests := []struct {
		input string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 20", 25},
		{"5 + 3 + 5 - 20", -7},
		{"-1000 + 400 - 20", -620},
		{"2 * 4 * 8 * 16", 1024},
		{"2 * 2 + 3 * 6", 22},
		{"-8 + 3 * 100 - 6 / 2", 289},
		{"(2 + 4) * 3 + 5 - ((2 + 6) * 4) / 2", 7},
		{"-20 * -2 * (10 + 4) / -2", -280},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true

}

func TestEvalBooleanExpressions(t *testing.T) {
	tests := []struct{
		input string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

// testBooleanObject - checks if boolean object matches expected value
// if not matches - throws an error with explanations
func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj is not boolean, got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, expected=%t, want=%t",
		expected, result.Value)
		return false
	}
	return true
}

// TestBangOperator - checks prefix operator `!`
// This operator inverts boolean variable (!true=false, !false=true)
// any integer with this prefix transforms into false
func TestBangOperator(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}