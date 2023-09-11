package piscine

func Capitalize(s string) string {
	var result string
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			if i == 0 {
				result += string(v - 32)
			} else if s[i-1] >= 'a' && s[i-1] <= 'z' || s[i-1] >= 'A' && s[i-1] <= 'Z' || s[i-1] >= '0' && s[i-1] <= '9' {
				result += string(v)
			} else {
				result += string(v - 32)
			}
		} else if v >= 'A' && v <= 'Z' {
			if i == 0 {
				result += string(v)
			} else if s[i-1] >= 'a' && s[i-1] <= 'z' || s[i-1] >= 'A' && s[i-1] <= 'Z' || s[i-1] >= '0' && s[i-1] <= '9' {
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
