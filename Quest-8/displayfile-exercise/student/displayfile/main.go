package main

// TODO: Complete the program to display file content
//
// This program should:
// 1. Check number of arguments
//    - 0 args: print "File name missing"
//    - 1 arg: read and display file content
//    - 2+ args: print "Too many arguments"
// 2. Open the file
// 3. Read the entire file
// 4. Print the content
//
// Hints:
// - Use os.Args[1:] to get arguments
// - Use len(args) to check argument count
// - Use os.Open(filename) to open file
// - Use io.ReadAll(file) to read content
// - Use defer file.Close() for cleanup
// - Handle errors by printing them with fmt.Println(err)

func main() {
	// args := os.Args[1:]

	// TODO: Check argument count and handle different cases
	// - If 0 arguments: print "File name missing" and return
	// - If more than 1 argument: print "Too many arguments" and return

	// TODO: Get the filename from args[0]

	// TODO: Open the file using os.Open()
	// Remember to check for errors!

	// TODO: Use defer to close the file

	// TODO: Read the file content using io.ReadAll()
	// Remember to check for errors!

	// TODO: Print the content as a string
	// Hint: Convert []byte to string, use fmt.Print (not Println)
}
