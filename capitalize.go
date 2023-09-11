package piscine

func IsLetterOrDigit(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	var result string
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			if i == 0 {
				result += string(v - 32)
			} else if IsLetterOrDigit(rune(s[i-1])) {
				result += string(v)
			} else {
				result += string(v - 32)
			}
		} else if v >= 'A' && v <= 'Z' {
			if i == 0 {
				result += string(v)
			} else if IsLetterOrDigit(rune(s[i-1])) {
				result += string(v + 32)
			} else {
				result += string(v)
			}
		} else if v >= '0' && v <= '9' {
			result += string(v)
		} else {
			result += string(v)
		}
	}
	return result
}
