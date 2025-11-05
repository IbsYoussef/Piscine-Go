package piscine

func AlphaCount(s string) int {
	count := 0

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') {
			count++
		}
	}

	return count
}
