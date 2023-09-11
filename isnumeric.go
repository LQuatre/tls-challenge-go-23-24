package piscine

func IsNumeric(str string) bool {
	for _, r := range str {
		if '0' > r || r > '9' {
			return false
		}
	}
	return true
}
