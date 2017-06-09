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
		{"1 < 2", true},
		{"5 < 1", false},
		{"3 > 5", false},
		{"100 > 44", true},
		{"3 == 3", true},
		{"5 == 30", false},
		{"3 != -100", true},
		{"-5 == -5", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"false == true", false},
		{"true != false", true},
		{"false != true", true},
		{"true != true", false},
		{"false != false", false},
		{"(1 < 10) == true", true},
		{"(1 < 10) != true", false},
		{"(1 > 10) != true", true},
		{"(1 > 10) != false", false},

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

// TestIfElseExpression - checks conditionals
func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input string
		expected interface{}
	}{
		{"if (true) { return 10; }", 10},
		{"if (false) { return 20; }", nil},
		{"if (1) { return 10; }", 10},
		{"if (10 < 20) { return 100; }", 100},
		{"if (10 > 20) { true }", nil},
		{"if (2 > 1) { return 1; } else { return 2; }", 1},
		{"if (2 < 1) { return 1; } else { return 2; }", 2},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

// testNullObject - checks if object is NULL object
func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL, got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 8;", 10},
		{"return 2 * 5; 8;", 10},
		{"100; return 2 * 5; 8;", 10},
		//{`
		//	if (10 > 1) {
		//		if (10 > 1) {
		//			return 10;
		//		}
		//		return 1;
		//	}
		//`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true;",
			"unknown operator: -BOOLEAN",
		},
		{
			"false + true;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"100; false + true; 25;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 2) { false + true; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (10 < 20) {
				if (20 < 30) {
					return false + true;
				}
				return 1;
			}
			return 0;
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
			evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
			tt.expectedMessage, errObj.Message)
		}
	}
}