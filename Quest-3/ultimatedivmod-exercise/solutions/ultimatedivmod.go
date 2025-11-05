package solutions

// Function Breakdown:
//
// UltimateDivMod(a *int, b *int):
//   - Takes two pointers to int
//   - Divides *a by *b
//   - Stores division result in *a
//   - Stores remainder in *b
//
// Approach:
//   - Read both values first (important!)
//   - Calculate division and modulo
//   - Store results back to pointers
//   - Use temp variables to avoid data loss
//
// Key Concepts:
//   - Pointers: Must dereference to read/write
//   - Order matters: Read before writing
//   - Division: a / b gives quotient
//   - Modulo: a % b gives remainder
//
// Example:
//   a := 13, b := 2
//   UltimateDivMod(&a, &b)
//   // a = 6, b = 1

func UltimateDivMod(a *int, b *int) {
	divResult := *a / *b // Calculate division (read first)
	modResult := *a % *b // Calculate remainder (read first)

	*a = divResult // Store division in a
	*b = modResult // Store remainder in b
}
