[![Build Status](https://travis-ci.org/technoboom/compiler.svg?branch=master)](https://travis-ci.org/technoboom/compiler)
# Building a compiler in Go programming language
This project was developed for building a compiler for own programming language (Beaver language).
## Done features
#### Lexer:
* Parses all the input until the end of the input
* Can recognize identifiers: e.g. `hello_world`
* Ignores whitespaces
* Added numbers (integers only): e.g. `246`
* Added keywords: `let`, `function`
* Added operators: `+`, `-`, `*`, `/`, `<`, `>`, `|`
* Added comparison and logical operators: `==`, `!=`
* Brackets and parenthesis support: `{}`, `()`
* Added keywords: `if`, `else`, `return`
* Recognizes comma and semicolon: `,`, `;`
* Recognizes EOF
* All not recognized symbols are ILLEGAL tokens: e.g. `$`

#### REPL for interpreter
* Uses Beaver lexer and prints given tokens from the input
Sample of REPL:
```
MacBook-Pro:compiler myuser$ go run main.go
Hello myuser! Starting Beaver REPL...
REPL (use Beaver commands):
beaver>>let number = 10;
{Type:LET Literal:let}
{Type:IDENT Literal:number}
{Type:= Literal:=}
{Type:INT Literal:10}
{Type:; Literal:;}
```

## Planned features
* C-like syntax
* Some elements of functional programming: closures, passing functions as arguments, returning function from function, assigning functions to variables
* Types: integer, boolean
* Conditions

## Limitations
* Only ASCII support

## How to improve:
* Add UTF-8 support
* Add new types: float, double, string, etc.
* Add new operators and operations
* Consider the space as token

## Getting started:
### Prerequisites
To install this software, you need to install Go programming language
(depends on OS you have): https://golang.org/doc/install  

[Not now] Or you can install this using Docker. In this case read the docs and installation
guides
### Installation
Clone the project:
```
git clone https://github.com/technoboom/compiler
```

### Run
Run the project:
```
go run ./main.go
```

### Testing
To run all tests in the project use command: `go test ./...`  
To test only lexer: `go test ./lexer`

## Quick intro into Beaver language:
### Syntax:
Beaver language supports C-like syntax.
Basic rules:
* all spaces ignored (maybe will be improved in the future)
* each sentence should contain semicolon at the end of the line

### Variables:
You can define new variable using `let` keyword.
```
let x = 10;
let y = true;
```
You can use English letters and underscore inside variable identifiers
```
let arabica_coffee = 95;
let _strength_percent = 50;
```

### Functions
The keyword `function` used for defining functions.
```
let multiply = function(a, b) {
    a * b;
}
```
Each function returns the last executed sentence.
In the sample above, the result of multiplication will be returned.

### Operators
You can use equal operator `==` for comparing two variables/values with one type:  
```
if (number == 1000000) {
    // perform the action to congratulate the one millionth visitor
}
```
Not equal operator `!=` gives you a way to get around:  
```
if (column != 1) {
    // ignore the column
}
```

## Intro into building compiler/interpreter:
Whether you are building an interpreter or a compiler most of the steps remain the same. The most common, basic steps are:
1. Lexical Analysis
2. Parsing
3. Semantic Analysis
4. Optimization
5. Code Generation

### 1. Lexical Analysis
Code - representation of commands for computers which is most suitable for human reading and writing.  
First step of building a compiler - performing lexical analysis.
Lexical analysis - process of scanning and splitting the code into small independent pieces - tokens.  
Each token is associated with a literal string (lexeme) that will be used in next steps.  
Literals (e.g., strings, integers, float numbers), keywords, operators are the main goals to recognize
on this step.

### 2. Parsing
During this step we are going to give some meanings to the tokens we received on the state of Lexical Analysis.  
Each token is an object and it's placed into a tree data structure.
On this step we need to take care about correct language syntax.   For different languages there are list of base rules: tabulation, opening and closing brackets, etc.

### 3. Semantic Analysis
On this stage we need to take care about correct language semantics.  
As an example, we need to ensure that when we have some variable with some type and we are going to assign another type to this variable we will get an error.

### 4. Optimization

### 5. Code Generation

## Resources:
1. Writting an interpreter in Go (Thorsten Ball): https://interpreterbook.com/
