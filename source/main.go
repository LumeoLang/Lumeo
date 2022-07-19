package main

import (
	"fmt"
	"lumeo/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		ews_error(err)
	}
	fmt.Printf("Lumeo running %s\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

func ews_error(err error) {
	fmt.Printf("ERROR: %s\n", err)
}
