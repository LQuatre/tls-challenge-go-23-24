package piscine

func ActiveBits(n int) int {
	ans := 0
	for n > 0 {
		if n%2 == 1 {
			ans++
		}
		n = n / 2
	}
	return ans
}
