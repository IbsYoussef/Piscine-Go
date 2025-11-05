package piscine

func Compare(a, b string) int {
	if a == "" && b == "" {
		return 0
	}
	if a == "" && b != "" {
		return -1
	}
	if a != "" && b == "" {
		return 1
	}

	aRunes := []rune(a)
	bRunes := []rune(b)

	for i := 0; i < len(aRunes) && i < len(bRunes); i++ {
		if aRunes[i] > bRunes[i] {
			return 1
		}
		if aRunes[i] < bRunes[i] {
			return -1
		}
	}

	if len(aRunes) > len(bRunes) {
		return 1
	}

	if len(aRunes) < len(bRunes) {
		return -1
	}

	return 0
}
