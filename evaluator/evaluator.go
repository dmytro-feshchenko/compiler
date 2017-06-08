package evaluator

import (
	"github.com/technoboom/compiler/ast"
	"github.com/technoboom/compiler/object"
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
		return &object.Boolean{Value: node.Value}
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