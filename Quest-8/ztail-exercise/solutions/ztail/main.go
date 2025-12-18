package main

import (
	"fmt"
	"io"
	"os"
)

// Status: Bonus

func main() {
	// Check minimum arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . -c <number> <file1> <file2> ...")
		os.Exit(1)
	}

	// Check -c flag
	if os.Args[1] != "-c" {
		fmt.Println("First argument must be -c")
		os.Exit(1)
	}

	// Convert number to int
	nBytes, err := atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number:", os.Args[2])
		os.Exit(1)
	}

	files := os.Args[3:]
	hasError := false

	// Process each file
	for i, filename := range files {
		// Print header if multiple files
		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", filename)
		}

		// Open file
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", filename)
			hasError = true
			continue
		}

		// Get file size
		info, err := file.Stat()
		if err != nil {
			fmt.Printf("error reading %s: %v\n", filename, err)
			file.Close()
			hasError = true
			continue
		}

		size := info.Size()

		// Calculate where to start reading
		start := size - int64(nBytes)
		if start < 0 {
			start = 0
		}

		// Seek to position
		_, err = file.Seek(start, 0)
		if err != nil {
			fmt.Printf("error seeking %s: %v\n", filename, err)
			file.Close()
			hasError = true
			continue
		}

		// Read to end of file
		content, err := io.ReadAll(file)
		file.Close()

		if err != nil {
			fmt.Printf("error reading %s: %v\n", filename, err)
			hasError = true
			continue
		}

		// Print content
		fmt.Print(string(content))
	}

	if hasError {
		os.Exit(1)
	}
}

// atoi converts string to int manually
func atoi(s string) (int, error) {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid digit")
		}
		result = result*10 + int(c-'0')
	}
	return result, nil
}
