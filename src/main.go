package main

import (
	"fmt"
	"os"

	"github.com/Urvirith/bearlang/src/repl"
)

func main() {
	fmt.Printf("This is the REPL of BearLang\n")
	fmt.Printf("Type in a command\n")
	repl.Start(os.Stdin, os.Stdout)
}
