package main

import (
	"os"

	"github.com/01-edu/z01"
)

// comcheck checks command-line arguments for specific strings
// and prints "Alert!!!" if any match is found
//
// Matches: "01", "galaxy", or "galaxy 01"
//
// Algorithm:
// 1. Get command-line arguments
// 2. Loop through each argument
// 3. Check if it matches any target string
// 4. If match found, print "Alert!!!" and exit
// 5. If no match, print nothing

func main() {
	args := os.Args[1:]

	// Check each argument
	for _, word := range args {
		if word == "01" || word == "galaxy" || word == "galaxy 01" {
			// Print "Alert!!!"
			for _, char := range "Alert!!!" {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
			return // Exit after first match
		}
	}
	// No match found - print nothing
}
