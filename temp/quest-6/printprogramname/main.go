package main

import (
	// for debugging
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fullPath := os.Args[0]
	lastSlashIndex := -1

	// Find the last slash
	for i, r := range fullPath {
		if r == '/' {
			lastSlashIndex = i
		}
	}

	// Get the program name
	var programName string
	if lastSlashIndex != -1 {
		programName = fullPath[lastSlashIndex+1:]
	} else {
		programName = fullPath
	}

	// Print each rune in the program name
	for _, v := range programName {
		if err := z01.PrintRune(v); err != nil {
			os.Stderr.WriteString("Failed to print the program name: " + programName + "\n")
			os.Exit(1)
		}
	}
	z01.PrintRune('\n') // Ensure the output ends with a newline
}
