package piscine

func myItoa(n int) string {
	var result string
	var temp string
	if n == 0 {
		return "0"
	}
	for n != 0 {
		temp = string(n%10 + 48)
		result = temp + result
		n /= 10
	}
	return result
}
