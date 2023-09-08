package piscine

func BasicAtoi(s string) int {
	arrayStr := []rune(s)
	n := len(s)
	num := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return num
		} else {
			num *= 10
			num += int(arrayStr[i]) - '0'
		}
	}
	return num
}
