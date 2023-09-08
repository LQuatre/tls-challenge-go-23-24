package piscine

func charToInt(r rune) int {
	var n int
	for i := '0'; i < r; i++ {
		n++
	}
	return n
}

func containsInt(r rune) bool {
	for i := '0'; i <= '9'; i++ {
		if r == i {
			return true
		}
	}
	return false
}

func Atoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}
	arrayStr := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if StrLen(s) < 1 {
			return 0
		}
	}
	n := 0
	for _, r := range s {
		if !containsInt(r) {
			return 0
		}
		n = n*10 + charToInt(r)
	}
	if arrayStr[0] == '-' {
		n *= -1
	}
	return n
}
