package student

// TODO: Write your code here
// Implement the PrintNbrBase function
//
// Hints:
// - Validate the base first
// - Base must have at least 2 characters
// - All characters must be unique
// - Cannot contain '+' or '-'
// - If invalid, print "NV" and return
//
// - Handle negative numbers: print '-', convert absolute value
// - Handle zero: print first character of base
//
// - Conversion algorithm:
//   1. Use modulo to get remainder (digit position in base)
//   2. Use division to reduce number
//   3. Build result in reverse order
//   4. Print result forwards
//
// - Consider creating a helper function to validate base
//
// Examples:
//   PrintNbrBase(125, "0123456789") prints "125"
//   PrintNbrBase(-125, "01") prints "-1111101"
//   PrintNbrBase(125, "0123456789ABCDEF") prints "7D"
//   PrintNbrBase(-125, "choumi") prints "-uoi"
//   PrintNbrBase(125, "aa") prints "NV"

func PrintNbrBase(nbr int, base string) {
	// Your code here
}
