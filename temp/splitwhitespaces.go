package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	// Add the last word if there is one
	if word != "" {
		result = append(result, word)
	}

	return result
}
