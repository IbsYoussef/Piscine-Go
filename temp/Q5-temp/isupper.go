package piscine

func IsUpper(s string) bool {
	if s == "" {
		return false
	}

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if !(runes[i] >= 'A' && runes[i] <= 'Z') {
			return false
		}
	}

	return true
}
