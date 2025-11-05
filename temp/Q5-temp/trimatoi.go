package piscine

func TrimAtoi(s string) int {
	result := 0
	runes := []rune(s)
	sign := 1
	foundSign := false

	for i := 0; i < len(runes); i++ {
		if runes[i] == '-' && !foundSign {
			sign = -1
		}
		if runes[i] >= '0' && runes[i] <= '9' {
			digit := int(runes[i] - '0')
			result = result*10 + digit
			foundSign = true
		}
	}

	return result * sign
}
