package main

import (
	"fmt"
	"io"
	"os"
)

// Program Breakdown:
//
// Goal: Display content of a file given as command-line argument
//
// Algorithm:
// 1. Check number of arguments
//    - 0 args: print "File name missing"
//    - 1 arg: read and display file
//    - 2+ args: print "Too many arguments"
// 2. Open file using os.Open()
// 3. Read entire file using io.ReadAll()
// 4. Print content to stdout
// 5. Handle errors gracefully
//
// Key Concepts:
// - os.Args: Command-line arguments
// - os.Open(): Open file for reading
// - io.ReadAll(): Read all content from file
// - defer: Ensure cleanup happens
// - Error handling: Check and handle errors
//
// Examples:
//   go run . → "File name missing"
//   go run . file1 file2 → "Too many arguments"
//   go run . test.txt → prints content of test.txt

// Status: Required

func main() {
	// Get arguments (excluding program name)
	args := os.Args[1:]

	// Check for no arguments
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	// Check for too many arguments
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	// Get filename
	filename := args[0]

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // Ensure file is closed when function exits

	// Read entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print content
	fmt.Print(string(content))
}
