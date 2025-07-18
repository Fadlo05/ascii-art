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
		fmt.Println("empty string")
		return
	}
	asciiMap := functions.ReadAsciiBanner("standard.txt")
	functions.AsciiRepresentation(input, asciiMap)
}
