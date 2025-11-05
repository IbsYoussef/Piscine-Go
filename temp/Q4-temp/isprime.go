package piscine

// Status: Optional

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	sqrt := 1
	for sqrt*sqrt <= nb {
		sqrt++
	}

	for i := 2; i < sqrt; i++ {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
