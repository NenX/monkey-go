package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		arg1 := args[1]
		err := repl.RunFile(arg1)
		if err != nil {
			fmt.Printf("Error running file: %s\n", err)
		}
		return
	}
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
