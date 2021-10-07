package main

import (
	"fmt"
	"os"

	"github.com/farbodahm/lets-go/monkeyLangInterpreter/repl"
)

func main() {
	fmt.Println("Monkey Lang interpreter...")
	fmt.Println("Write your Monkey lang statements.")
	repl.Start(os.Stdin, os.Stdout)
}
