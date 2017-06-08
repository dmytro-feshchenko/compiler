package evaluator

import (
	"github.com/technoboom/compiler/ast"
	"github.com/technoboom/compiler/object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value:true}
	FALSE = &object.Boolean{Value:false}
)

// Eval - evaluates current node (traverses AST)
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}
	return nil
}

// evalStatements - evaluates array of statements in a while
// returns result of evaluation
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

// nativeBoolToBooleanObject - transforms bool variable into object.Object
// structure. This method used to increase performance (because we won't need
// to manually transform bool to object each time we need this)
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// evalPrefixExpression - contains switch to decide how to evaluate expression
// with prefix structure
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
}

// evalBangOperatorExpression - evaluates prefix expression with `!` operator
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}