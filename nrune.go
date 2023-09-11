package piscine

func NRune(s string, n int) rune {
	result := []rune(s)
	nb := n - 1
	if nb > len(s) {
		return 0
	}
	return result[nb]
}
