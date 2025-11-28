package student

// TODO: Write your code here
// Convert number from one base to another
//
// This combines AtoiBase and PrintNbrBase!
//
// Algorithm:
// 1. Convert nbr from baseFrom to decimal integer
//    - Use Horner's method like AtoiBase
//    - result = result * len(baseFrom) + digitValue
//
// 2. Convert decimal to baseTo string
//    - Use modulo/division like PrintNbrBase
//    - remainder = n % len(baseTo)
//    - Build result string
//
// Helper functions you'll need:
// - atoiBase(s, base string) int
// - printNbrBase(n int, base string) string
// - indexOf(s string, char rune) int
//
// Examples:
//   ConvertBase("101011", "01", "0123456789")
//   Binary "101011" → Decimal 43 → Decimal "43"
//   Returns: "43"
//
//   ConvertBase("2A", "0123456789ABCDEF", "01")
//   Hex "2A" → Decimal 42 → Binary "101010"
//   Returns: "101010"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Your code here
	return ""
}
