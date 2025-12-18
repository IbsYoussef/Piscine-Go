package main

// TODO: Implement simplified tail -c command
//
// Steps:
// 1. Check arguments (need at least 4)
// 2. Validate -c flag
// 3. Convert number string to int
// 4. Loop through files:
//    - Print header if multiple files
//    - Open file
//    - Get size with file.Stat()
//    - Calculate start position: size - N
//    - Seek to start position
//    - Read rest of file with io.ReadAll()
//    - Print content
// 5. Exit with 1 if any errors

func main() {
	// TODO: Implement the program here
}

// TODO: Create atoi function to convert string to int
// func atoi(s string) (int, error) { ... }
