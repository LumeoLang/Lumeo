package main

import (
	"log"
	"lumeo/repl"
	"os"
	"strings"
)

func main() {
	path, err := os.Getwd()
	path = path + "\\basic.lmo"
	if err != nil {
		log.Println(err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(string(content))
	repl.Start(r, os.Stdout)
}
