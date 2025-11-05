package piscine

func Capitalize(s string) string {
	inWord := false
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' && !inWord {
			// Capitalize the first lowercase letter of a word
			runes[i] -= 32
			inWord = true
		} else if runes[i] >= 'A' && runes[i] <= 'Z' && inWord {
			// Lowercase subsequent uppercase letters within a word
			runes[i] += 32
		} else if runes[i] >= '0' && runes[i] <= '9' {
			inWord = true
		} else {
			// Reset inWord after any non-letter character, including numbers
			inWord = (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z')
		}
	}
	return string(runes)
}
