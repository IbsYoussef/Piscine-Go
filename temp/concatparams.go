package piscine

func ConcatParams(args []string) string {
	result := ""

	for i, v := range args {
		result += v
		if i < len(args)-1 {
			result += "\n"
		}
	}

	return result
}
