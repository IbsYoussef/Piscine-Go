package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for correct usage
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . -c <number_of_characters> <file1> <file2> ...")
		os.Exit(1)
	}

	// Ensure the first argument is the '-c' flag
	if os.Args[1] != "-c" {
		fmt.Println("The first argument must be '-c'")
		os.Exit(1)
	}

	// Convert the number of characters to an integer
	nChars, err := asciiToInt(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number of characters:", os.Args[2])
		os.Exit(1)
	}

	// Get the list of files to process
	files := os.Args[3:]

	// Process each file
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("open %s: %v\n", filename, err)
			os.Exit(1) // Exit with an error if any file fails to open
		}
		defer file.Close()

		// Read the file content and get the last `nChars` characters
		content, err := readLastNChars(file, nChars)
		if err != nil {
			fmt.Printf("error reading file %s: %v\n", filename, err)
			os.Exit(1)
		}

		// If there are multiple files, print the file name header
		if len(files) > 1 {
			fmt.Printf("\n==> %s <==\n", filename)
		}
		fmt.Print(content)
	}
}

// asciiToInt converts a string of digits to an integer.
func asciiToInt(s string) (int, error) {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, fmt.Errorf("invalid character in number: %c", s[i])
		}
		result = result*10 + int(s[i]-'0')
	}
	return result, nil
}

// readLastNChars reads the last `n` characters of a file.
func readLastNChars(file *os.File, n int) (string, error) {
	// Get the file size
	fi, err := file.Stat()
	if err != nil {
		return "", err
	}

	// Calculate the offset to start reading
	offset := fi.Size() - int64(n)
	if offset < 0 {
		offset = 0
	}

	// Move the file pointer to the offset
	_, err = file.Seek(offset, 0)
	if err != nil {
		return "", err
	}

	// Read the last `n` characters
	buf := make([]byte, n)
	_, err = file.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
