package parser

import (
	"fmt"
	"strings"
)

// contains current trace level for printing messages with indentation
var traceLevel int = 0

// character which will be used for showing indentation level
const traceIndentationPlaceholder string = "\t"

// indentationLevel - returns current indentation level
// as string which contains some number of indentation placeholders
// according to current traceLevel
func indentationLevel() string {
	return strings.Repeat(traceIndentationPlaceholder, traceLevel-1)
}

// tracePrint - prints message with current indentation level using fmt
func tracePrint(fs string) {
	fmt.Printf("%s%s\n", indentationLevel(), fs)
}

// increaseIndentation - increases current traceLevel
func increaseIndentation() { traceLevel = traceLevel + 1 }
// decreaseIndentation - decreases current traceLevel
func decreaseIndentation() { traceLevel = traceLevel - 1 }

// trace - increases indent prints begin message
func trace(msg string) string {
	increaseIndentation()
	tracePrint("BEGIN " + msg)
	return msg
}

// untrace - prints end message and decreases indentation
func untrace(msg string) {
	tracePrint("END " + msg)
	decreaseIndentation()
}
