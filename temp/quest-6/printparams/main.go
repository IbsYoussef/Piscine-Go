package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, v := range args {
		for _, r := range v {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
