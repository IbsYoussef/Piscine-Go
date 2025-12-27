package student

// TODO: Implement Itoa function
// Convert integer to string (reverse of Atoi)
//
// Algorithm:
// 1. Handle sign:
//    - If n < 0: save sign as "-" and make n positive
// 2. Handle zero edge case: return "0"
// 3. Extract digits using modulo/division:
//    - digit = n % 10 (gets last digit)
//    - Convert digit to char: '0' + digit
//    - Prepend to result string
//    - n = n / 10 (remove last digit)
// 4. Prepend sign to result
//
// Key conversion:
//   digit 5 → '0' + 5 = '5' (character)
//
// Example for 1234:
//   Extract 4 → "4"
//   Extract 3 → "34"
//   Extract 2 → "234"
//   Extract 1 → "1234"

func Itoa(n int) string {
	// Your code here
	return ""
}
