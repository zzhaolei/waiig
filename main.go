package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/zzhaolei/waiig/repl"
)

func main() {
	user2, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user2.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
