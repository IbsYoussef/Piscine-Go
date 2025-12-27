package solutions

// WordCount struct to hold word and its count
type WordCount struct {
	word  string
	count int
}

// MaxWordCountN returns the top n most frequent words
//
// Algorithm:
// 1. Split text into words (manual splitting by spaces/newlines/tabs)
// 2. Count word frequencies using a map
// 3. Convert map to slice for sorting
// 4. Sort by count (descending), then by word (ascending if tied)
// 5. Take top n and return as map
//
// Sorting criteria:
// - Primary: Higher count first
// - Secondary: Lower ASCII word first (if counts equal)

func MaxWordCountN(text string, n int) map[string]int {
	// Step 1: Split into words
	words := splitWords(text)

	// Step 2: Count frequencies
	freq := countWords(words)

	// Step 3: Convert to slice
	slice := mapToSlice(freq)

	// Step 4: Sort
	sortWordCounts(slice)

	// Step 5: Get top n
	return topN(slice, n)
}

// splitWords splits text by spaces, newlines, and tabs
func splitWords(text string) []string {
	words := []string{}
	word := ""

	for _, char := range text {
		if char == ' ' || char == '\n' || char == '\t' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	// Don't forget last word
	if word != "" {
		words = append(words, word)
	}

	return words
}

// countWords counts frequency of each word
func countWords(words []string) map[string]int {
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

// mapToSlice converts map to slice for sorting
func mapToSlice(freq map[string]int) []WordCount {
	slice := []WordCount{}
	for word, count := range freq {
		slice = append(slice, WordCount{word, count})
	}
	return slice
}

// sortWordCounts sorts by count (desc), then by word (asc)
func sortWordCounts(slice []WordCount) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			// Sort by count descending
			if slice[j].count < slice[j+1].count {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			} else if slice[j].count == slice[j+1].count {
				// If equal, sort by word ascending
				if slice[j].word > slice[j+1].word {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			}
		}
	}
}

// topN returns map of top n words
func topN(slice []WordCount, n int) map[string]int {
	result := make(map[string]int)

	count := n
	if len(slice) < n {
		count = len(slice)
	}

	for i := 0; i < count; i++ {
		result[slice[i].word] = slice[i].count
	}

	return result
}
