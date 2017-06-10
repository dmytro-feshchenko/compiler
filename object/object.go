package object

import (
	"fmt"
	"github.com/technoboom/compiler/ast"
	"bytes"
	"strings"
)

// ObjectType - represents type of object
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ = "STRING"
)

// Object - interface for representing types objects
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer - represents int type
type Integer struct {
	Value int64
}

// Inspect - shows value of the object
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type - returns type of the object
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Boolean - represent bool variables
type Boolean struct {
	Value bool
}

// Inspect - shows value of the object
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type - returns type of the object
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Null - Represents null type
type Null struct {}

// Inspect - shows value of the object
func (n *Null) Inspect() string {
	return "null"
}

// Type - returns type of the object
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// ReturnValue - object that contains return value of function
type ReturnValue struct {
	Value Object
}

// Type - returns type of the object
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect - shows value of the object
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error - structure that stores error messages for error handling
type Error struct {
	Message string
}

// Type - returns type of the object
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect - shows value of the object
func (e *Error) Inspect() string {
	return e.Message
}

// Function - represents function structure
type Function struct {
	// parameters of the function
	Parameters []*ast.Identifier
	// statements inside block statement of the function
	Body *ast.BlockStatement
	// environment variables
	Env *Environment
}

// Type - returns type of the object
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// Inspect - shows value of the object
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("function")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String - represents string type
type String struct {
	Value string
}

// Type - returns type of the object
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect - shows value of the object
func (s *String) Inspect() string {
	return s.Value
}