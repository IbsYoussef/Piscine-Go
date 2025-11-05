package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if !(runes[i] >= 'A' && runes[i] <= 'Z') &&
			!(runes[i] >= 'a' && runes[i] <= 'z') &&
			!(runes[i] >= '0' && runes[i] <= '9') {
			return false
		}
	}

	return true
}
