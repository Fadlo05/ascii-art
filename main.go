package main

import (
	"fmt"
	"os"

	"asciiArt/functions"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid input!")
		return
	}
	input := args[0]
	if len(input) == 0 {
		fmt.Println("empy string")
		return
	}
	functions.ReadAsciiBanner("standard.txt")
}
