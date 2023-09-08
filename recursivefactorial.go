package piscine

func RecursiveFactorial(nb int) int {
	if nb != 0 {
		if nb == 1 {
			return 1
		} else {
			return nb * RecursiveFactorial(nb-1)
		}
	} else {
		return 0
	}
}
