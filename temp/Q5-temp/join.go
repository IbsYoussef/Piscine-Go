package piscine

func Join(strs []string, sep string) string {
	result := ""

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		result = strs[0]
		return result + sep
	}

	for i, v := range strs {
		if i > 0 {
			result += sep
		}
		result += v
	}

	return result
}
