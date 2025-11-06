package student

// TODO: Write your code here
// Implement BOTH the IsPrime helper function AND FindNextPrime function
//
// First, implement IsPrime (or copy from isprime exercise):
// - Function signature: func IsPrime(nb int) bool
// - Handle edge case: if nb <= 1 { return false }
// - Calculate sqrt: sqrt := 1; for sqrt*sqrt <= nb { sqrt++ }
// - Check divisors: for i := 2; i < sqrt; i++
// - If divisible: if nb%i == 0 { return false }
// - If no divisors: return true
//
// Then, implement FindNextPrime:
// - Function signature: func FindNextPrime(nb int) int
// - Handle edge case: if nb <= 1 { return 2 }
// - Check if nb is prime: if IsPrime(nb) { return nb }
// - Use infinite loop: for { ... }
// - Increment nb: nb++
// - Check if prime: if IsPrime(nb) { return nb }
//
// Note: You need BOTH functions in this file!
//
// Examples:
//   FindNextPrime(5) returns 5
//   FindNextPrime(4) returns 5
//   FindNextPrime(14) returns 17

// IsPrime helper function - implement this first
func IsPrime(nb int) bool {
	// Your code here
	return false
}

// FindNextPrime main function - implement this second
func FindNextPrime(nb int) int {
	// Your code here
	return 0
}
