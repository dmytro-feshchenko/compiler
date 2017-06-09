// Package repl - implements REPL for Beaver interpreter
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/technoboom/compiler/lexer"
	"github.com/technoboom/compiler/parser"
	"github.com/technoboom/compiler/evaluator"
	"github.com/technoboom/compiler/object"
)

// PROMPT - the message that leads each line of the REPL
const PROMPT = "beaver>>"

const BEAVER = `
     __________
    /  _    _  \
  _/   _    _   \_
 |_|  | |  | |  |_|
  \   |_|  |_|   /
   |      _     |
   |    | | |   |
   |            |
   |____________|

`

// Start - starts the REPL interactive mode
// reads from input until encountering new line,
// converts input into tokens and prints them all
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		// print the prompt
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			// stop the REPL
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// printParseErrors - prints parser errors
func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, BEAVER)
	io.WriteString(out, "Woops! Something got wrong here:")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}