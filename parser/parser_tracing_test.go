package parser

import (
	"testing"
	"strings"
)

func TestIncreaseIndentationLevel(t *testing.T) {
	for i := 0; i < 10; i++ {
		traceLevel = 0
		// increase indentation level i times
		for j := 0; j < i; j++ {
			increaseIndentation()
		}
		if traceLevel != i {
			t.Errorf("traceLevel after increaseIndendation() level wrong")
		}
	}
}

func TestDecreaseIndentationLevel(t *testing.T) {
	for i := 0; i < 10; i++ {
		traceLevel = 100
		// increase indentation level i times
		for j := 0; j < i; j++ {
			decreaseIndentation()
		}
		if traceLevel != 100 - i {
			t.Errorf("traceLevel after decreaseIndendation() level wrong")
		}
	}
}

func TestIndentationLevelString(t *testing.T) {

	for i := 0; i < 5; i++ {
		traceLevel = i + 1
		realIndentation := indentationLevel()
		if realIndentation != strings.Repeat("\t", i) {
			t.Errorf("indentationLevel() returns wrong string, got=%s",
				realIndentation)
		}
	}
}