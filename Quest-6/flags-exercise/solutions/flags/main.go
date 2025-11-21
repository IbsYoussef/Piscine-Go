package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Handle command-line flags for string manipulation
//
// Approach:
//   - Parse arguments to detect flags and values
//   - Extract insert value if --insert or -i present
//   - Set boolean flag if --order or -o present
//   - Get main string argument (non-flag)
//   - Apply operations: insert first, then order
//   - Print result or help message
//
// Key Concepts:
//   - Flag parsing: Detecting --flag=value and -f=value
//   - String slicing: Extracting values after "="
//   - Boolean flags: No value needed, just presence
//   - Help message: Formatted output with tabs and spaces
//   - Bubble sort: Sorting the string characters
//
// Flag Formats:
//   --insert=value or -i=value: Insert value into string
//   --order or -o: Sort the result in ASCII order
//   --help or -h: Print help message
//
// Operations Order:
//   1. Get main string
//   2. Append insert value (if flag present)
//   3. Sort result (if flag present)
//   4. Print final result
//
// Examples:
//   ./flags --insert=4321 --order asdad → 1234aadds
//   ./flags --insert=4321 asdad → asdad4321
//   ./flags --order 43a21 → 1234a
//   ./flags → prints help

// Status: Optional

func printhelp() {
	help := "--insert\n"
	help += "  -i\n"
	help += "\t This flag inserts the string into the string passed as argument.\n"
	help += "--order\n"
	help += "  -o\n"
	help += "\t This flag will behave like a boolean, if it is called it will order the argument.\n"

	for _, char := range help {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]
	insertValue := ""
	hasInsert := false
	hasOrder := false
	mainString := ""

	// Print help if no arguments
	if len(args) == 0 {
		printhelp()
		return
	}

	// Parse all arguments
	for _, arg := range args {
		// Check for help flag
		if arg == "--help" || arg == "-h" {
			printhelp()
			return
		}

		// Check for --insert=value
		if len(arg) > 9 && arg[:9] == "--insert=" {
			hasInsert = true
			insertValue = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			// Check for -i=value
			hasInsert = true
			insertValue = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			// Check for order flag (boolean)
			hasOrder = true
		} else if arg != "" && arg[0] != '-' {
			// Non-flag argument is the main string
			mainString = arg
		}
	}

	// Start with main string
	result := mainString

	// Apply insert operation
	if hasInsert {
		result += insertValue
	}

	// Apply order operation (bubble sort)
	if hasOrder {
		runeResult := []rune(result)
		for i := 0; i < len(runeResult)-1; i++ {
			for j := 0; j < len(runeResult)-i-1; j++ {
				if runeResult[j] > runeResult[j+1] {
					runeResult[j], runeResult[j+1] = runeResult[j+1], runeResult[j]
				}
			}
		}
		// Convert sorted runes back to string
		result = string(runeResult)
	}

	// Print final result
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
