package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	alphabet := make(map[int]rune)

	for i := 1; i <= 26; i++ {
		alphabet[i] = rune('a' + i - 1)
	}

	isUpper := false
	args := os.Args[1:]

	if len(args) == 0 {
		return // No arguments, exit without printing anything
	}

	if args[0] == "--upper" {
		isUpper = true
		args = args[1:] // Skip the flag in further processing
	}

	for _, arg := range args {
		n := 0
		valid := true

		// Manual conversion from string to number
		for _, r := range arg {
			if r < '0' || r > '9' {
				valid = false
				break
			}
			n = n*10 + int(r-'0')
		}

		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ') // Invalid input print space
		} else {
			letter := alphabet[n]
			if isUpper {
				letter = letter - 'a' + 'A' // Convert to uppercase
			}
			z01.PrintRune(letter)
		}
	}

	z01.PrintRune('\n')
}
