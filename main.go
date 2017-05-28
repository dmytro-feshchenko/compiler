package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/technoboom/compiler/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Starting Beaver REPL...\n", user.Username)
	fmt.Printf("REPL (use Beaver commands):\n")
	repl.Start(os.Stdin, os.Stdout)
}
