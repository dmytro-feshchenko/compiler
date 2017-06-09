package object

// NewEnclosedEnvironment - creates new enclosed environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment - creates new environment and returns it
func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s, outer: nil}
}

// Environment - stores variables objects
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get - returns variable from environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		// try to find in parent scope
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set - creates variable in environment if not exists
// otherwise replaces variable value
func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}