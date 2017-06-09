[![Build Status](https://travis-ci.org/technoboom/compiler.svg?branch=master)](https://travis-ci.org/technoboom/compiler)
[![codecov](https://codecov.io/gh/technoboom/compiler/branch/master/graph/badge.svg)](https://codecov.io/gh/technoboom/compiler)
[![Issue Count](https://codeclimate.com/github/technoboom/compiler/badges/issue_count.svg)](https://codeclimate.com/github/technoboom/compiler)

# Building a compiler in Go programming language
This project was developed for building a compiler for own programming language (Beaver language).  
Note: I use a lot of ideas and boilerplate from awesome book "Writing an interpreter in Go"(Thorsten Ball) to create my
own compiler. To see other resources I used within the project, see the <a href="#resources">resources section.</a>

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

#### Parser
We used "top down operator precedence" parser, also known as "Pratt parser"

##### Features
- [x] Parsing let expressions: `let a = 10;`
- [x] Parsing return statements: `return 500;`
- [x] Parsing int expressions: `5;`
- [x] Prefix operators: `<prefix operator><expression>;`, `-10;`, `!true;`  
Supports two operators: `!` and `-`
- [x] Infix operators `<expression><infix operator><expression>`, e.g. `5 + 10`, `2 - 8`  
Supports 8 operators: `+`, `-`, `*`, `/`, `<`, `>`, `==`, `!=`
- [x] Working with operations precedences
- [x] Parsing function literals: `function(x, y) {}`
- [x] Call expressions: `<expression>(<comma separated expressions>)`

##### Samples:
* `let` statement
```
let foo = 27;
let bar = 50;
let foobar = foo + bar;
```
* `return` statement
```
return foo;
return foo + bar;
return foobar(foo, bar);
```
* prefix operators
```
-5;
-100;
!true;
```
* infix operators
```
5 + 100;
5 - 6;
10 / 5;
2 * 90;
4 == 4;
3 != 6;
2 < 87;
5 > 4;
```

#### REPL for interpreter
* Uses Beaver parser and prints given tokens from the input

Sample of REPL:
```
MacBook-Pro:compiler myuser$ go run main.go
Hello myuser! Starting Beaver REPL...
REPL (use Beaver commands):
beaver>>1 + 100
101
beaver>>true == false
false
beaver>>(2 + 10) * (5 + 1)
72
beaver>>if (10 > 100) { true } else { false }
false
beaver>>if (true) { 100 / 5 }
20
beaver>>let x 12 * 3

     __________
    /  _    _  \
  _/   _    _   \_
 |_|  | |  | |  |_|
  \   |_|  |_|   /
   |      _     |
   |    | | |   |
   |            |
   |____________|

Woops! Something got wrong here:
expected next token to be '=', got 'INT' instead
```

#### Evaluator:
- [x] Just evaluates statements and expressions in a while
- [x] Can evaluate integers expressions
- [x] Can evaluate boolean expressions
- [x] Can evaluate null
- [x] Can evaluate prefix expressions: `!`, `-`
- [x] Can evaluate infix expressions for integers: `+`, `-`, `*`, `/`
- [x] Can evaluate infix expressions for comparing: `==`, `!=`, `>`, `<`
- [x] Can evaluate conditionals:  
`if (conditional) { consequence }` or `if (conditional) { consequence } else { alternative }`
- [x] Can evaluate return statements
- [x] Error Handling
- [ ] Can evaluate functions calls

### Types:
- [x] Integers
- [x] Booleans
- [x] Null

## Planned features
* C-like syntax
* Some elements of functional programming: closures, passing functions as arguments, returning function from function, assigning functions to variables
* Types: integer, boolean
* Conditions

## Limitations
* Only ASCII support

## How to improve:
* Add UTF-8 support (using the )
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

### Run REPL
Run the REPL:
```
go run ./main.go
```

### Testing
To run all tests in the project use command: `go test ./...`  
To test only lexer: `go test ./lexer`  
To test only parser: `go test ./parser`  
To test only ast: `go test ./ast`  
To test tokens: `go test ./tokens`

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
let camelCase = true;
let UpperCamelCase = false;
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

### Conditions
In Beaver we can use keywords `if` nad `else` to work with conditionals
```
if (temperature > 0) {
    // it's hot enough
} else {
    // you can mold snowballs
}
```

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
You can find my implementation of Lexer in `./lexer` folder.  
Also we use structures from `./token` inside our parser.

### 2. Parsing
During this step we are going to give some meanings to the tokens we received on the state of Lexical Analysis.  
Each token is an object and it's placed into a tree data structure.
On this step we need to take care about correct language syntax.   For different languages there are list of base rules: tabulation, opening and closing brackets, etc.

#### Pratt Parser
You can find my implementation of Prett Parser in `./parser` folder.  
Also we use structures from `./ast` inside our parser.

### 3. Semantic Analysis
On this stage we need to take care about correct language semantics.  
As an example, we need to ensure that when we have some variable with some type and we are going to assign another type to this variable we will get an error.

### 4. Optimization
On this stage we need to think about performance of our application. To do this,
we should remove overhead constructions and operations.  
We can choose 1 of the options to perform this operation:
1. At the stage of processing intermediate code
2. At the stage of processing machine code (or other low-level representation)   

### 5. Code Generation
We use intermediate code to produce targeted low-level code.

## Resources:
1. Writing an interpreter in Go (Thorsten Ball)
2. Compiler Construction by Niklaus Wirth (2014)
3. Series of blog posts about building compilers:  http://noeffclue.blogspot.ca/2014/05/compiler-part-1-introduction-to-writing.html
