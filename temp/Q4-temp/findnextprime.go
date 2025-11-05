package piscine

// Status: Optional

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	if IsPrime(nb) {
		return nb
	}

	for {
		nb++
		if IsPrime(nb) {
			return nb
		}
	}
}
