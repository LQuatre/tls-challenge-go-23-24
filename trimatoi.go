package piscine

func TrimAtoi(s string) int {
	var result int
	var sign int = 1
	for _, v := range s {
		if v == '-' && result == 0 {
			sign = -1
		}
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		}
	}
	return result * sign
}
