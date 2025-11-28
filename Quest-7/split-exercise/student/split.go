package student

// TODO: Write your code here
// Split string by custom separator
//
// Algorithm:
// 1. Initialize: result := []string{}, word := ""
// 2. Handle edge case: if len(sep) == 0, return []string{s}
// 3. Loop through string with manual index (for i := 0; i < len(s); )
// 4. Check if separator found at current position:
//    - Use: s[i:i+len(sep)] == sep
//    - Don't forget to check: i+len(sep) <= len(s)
// 5. If separator found:
//    - Save word: result = append(result, word)
//    - Reset word: word = ""
//    - Skip separator: i += len(sep)
// 6. If no separator:
//    - Add character: word += string(s[i])
//    - Move forward: i++
// 7. After loop: append last word
//
// Key trick:
//   s[i:i+len(sep)] extracts a substring of length len(sep) starting at i
//   This lets you check multi-character separators!
//
// Examples:
//   Split("HelloHAhowHAareHAyou?", "HA")
//   Returns: []string{"Hello", "how", "are", "you?"}
//
//   Split("a,b,c", ",")
//   Returns: []string{"a", "b", "c"}

func Split(s, sep string) []string {
	// Your code here
	return nil
}
