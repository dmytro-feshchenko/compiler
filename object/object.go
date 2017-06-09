package object

import "fmt"

// ObjectType - represents type of object
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
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

// NewEnvironment - creates new environment and returns it
func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

// Environment - stores variables objects
type Environment struct {
	store map[string]Object
}

// Get - returns variable from environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set - creates variable in environment if not exists
// otherwise replaces variable value
func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}