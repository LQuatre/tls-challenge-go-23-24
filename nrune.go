package piscine

func NRune(s string, n int) rune {
	result := []rune(s)
	if n > len(s) || n < 0 {
		return 0
	}
	nb := n - 1
	return result[nb]
}
