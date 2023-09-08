package piscine

func charToInt(ch rune) int {
	var nm int
	for i := '0'; i < ch; i++ {
		nm++
	}
	return nm
}

func containsInt(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}
	return false
}

func Atoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}

	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if StrLen(s) < 1 {
			return 0
		}
	}

	nm := 0

	for _, ch := range s {
		if !containsInt(ch) {
			return 0
		}
		nm = nm*10 + charToInt(ch)
	}

	if s0[0] == '-' {
		nm *= -1
	}
	return nm
}
