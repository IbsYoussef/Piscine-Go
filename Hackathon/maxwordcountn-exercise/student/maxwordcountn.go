package student

// TODO: Implement MaxWordCountN function
// Return top n most frequent words
//
// High-level algorithm:
// 1. Split text into words (by spaces, newlines, tabs)
// 2. Count word frequencies using a map
// 3. Convert map to slice for sorting
// 4. Sort by count (highest first), then by word (alphabetically if tied)
// 5. Take top n and return as map
//
// Sorting priority:
// - Higher count comes first
// - If counts equal, lower ASCII word comes first
//
// Helper functions needed:
// - splitWords(text string) []string
// - countWords(words []string) map[string]int
// - sortWordCounts(slice []WordCount)
// - topN(slice []WordCount, n int) map[string]int

func MaxWordCountN(text string, n int) map[string]int {
	// Your code here
	return nil
}
