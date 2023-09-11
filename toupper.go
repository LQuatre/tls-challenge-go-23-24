package piscine

func ToUpper(s string) string {
	result := ""
	if IsUpper(s) {
		return s
	}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(r - 32)
		} else {
			result += string(r)
		}
	}
	return result
}
