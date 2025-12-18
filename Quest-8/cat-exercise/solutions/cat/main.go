package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Get command-line arguments (excluding program name)
	args := os.Args[1:]

	// MODE 1: No arguments - read from stdin and echo to stdout
	if len(args) == 0 {
		// Copy all data from stdin to stdout (creates echo effect)
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// MODE 2: Files provided - read and print each file
	hasError := false

	// Process each file
	for _, filename := range args {
		// Try to open the file
		file, err := os.Open(filename)
		if err != nil {
			// Print error to stderr and continue with next file
			fmt.Fprintf(os.Stderr, "ERROR: open %s: no such file or directory\n", filename)
			hasError = true
			continue
		}

		// Copy file contents to stdout
		_, err = io.Copy(os.Stdout, file)
		file.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: reading %s: %v\n", filename, err)
			hasError = true
		}
	}

	// Exit with error status if any errors occurred
	if hasError {
		os.Exit(1)
	}
}
