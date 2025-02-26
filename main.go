package main

import (
	"fmt"
	"meze/waiig/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
