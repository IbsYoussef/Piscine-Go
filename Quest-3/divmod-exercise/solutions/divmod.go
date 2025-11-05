package solutions

// Function Breakdown:
//
// DivMod(a int, b int, div *int, mod *int):
//   - Takes two integers to divide
//   - Takes two pointers to store results
//   - Stores division result in *div
//   - Stores remainder in *mod
//
// Approach:
//   - Calculate division: a / b
//   - Calculate modulo: a % b
//   - Store results at pointer addresses
//
// Key Concepts:
//   - Division operator (/): Integer division
//   - Modulo operator (%): Remainder after division
//   - Pointer assignment: *ptr = value
//
// Example:
//   a := 13, b := 2
//   var div, mod int
//   DivMod(a, b, &div, &mod)
//   // div = 6, mod = 1

func DivMod(a int, b int, div *int, mod *int) {
	divResult := a / b // Calculate division
	modResult := a % b // Calculate remainder
	*div = divResult   // Store division result
	*mod = modResult   // Store remainder
}
