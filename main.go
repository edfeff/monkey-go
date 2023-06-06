package main

import (
	"fmt"
	"monkey/explainer"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) == 1 {
		startWithRepl()
	}
	args := os.Args
	starWithFile(args)
}

func starWithFile(args []string) {
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	explainer.Start(file, os.Stdout)
}

func startWithRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to types in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
