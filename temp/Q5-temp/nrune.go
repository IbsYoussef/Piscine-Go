package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)

	if n <= 0 || n > len(runes) {
		return rune(0)
	}

	return runes[n-1]
}
