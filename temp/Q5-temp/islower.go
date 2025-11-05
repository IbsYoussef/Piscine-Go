package piscine

func IsLower(s string) bool {
	if s == "" {
		return false
	}

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if !(runes[i] >= 'a' && runes[i] <= 'z') {
			return false
		}
	}

	return true
}
