// Package repl - implements REPL for Beaver interpreter
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/technoboom/compiler/lexer"
	"github.com/technoboom/compiler/token"
)

// PROMPT - the message that leads each line of the REPL
const PROMPT = "beaver>>"

// Start - starts the REPL interactive mode
// reads from input until encountering new line,
// converts input into tokens and prints them all
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
