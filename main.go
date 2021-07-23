package main

import (
	"fmt"
	"go-interpreter-demo/repl"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! \n", current.Username)
	fmt.Printf("Please type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
