package piscine

func NRune(s string, n int) rune {
	result := []rune(s)
	if n > len(s) {
		return 0
	}
	nb := n - 1
	return result[nb]
}
