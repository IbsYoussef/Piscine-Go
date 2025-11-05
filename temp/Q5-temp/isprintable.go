package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)

	for _, r := range runes {
		if !(r >= 32 && r <= 126) {
			return false
		}
	}

	return true
}
