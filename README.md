# Building a compiler in Go programming language
This project was developed for building a compiler for own programming language.
## Done features
[Still in development]
## Planned features
* C-like syntax
* Some elements of functional programming: closures, passing functions as arguments, returning function
* Types: integer, boolean
* Conditions
## Limitations
* ASCII support
## How to improve:
* Add UTF-8 support
* Add new types
* Add new operators and operations
## Getting started:
### Prerequisites
To install this software, you need to install Go programming language
(depends on OS you have): https://golang.org/doc/install
[Not now] Or you can install this using Docker. In this case read the docs and installation
guides
### Installation
Clone the project:
```
git clone github.com/technoboom/compiler
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
On this step we need to take care about correct language syntax. For different languages there are list of base rules: tabulation, opening and closing brackets, etc.

### 3. Semantic Analysis
On this stage we need to take care about correct language semantics. As an example, we need to ensure that when we have some variable with some type and we are going to assign another type to this variable we will get an error.

### 4. Optimization

### 5. Code Generation
