package main

import (
	"fmt"
	"lumeo/repl"
	"os"
)

func main() {
	fmt.Println("This will be Lumeo console :)")
	repl.Start(os.Stdin, os.Stdout)
}
