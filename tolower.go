package piscine

func ToLower(s string) string {
	result := ""
	if IsLower(s) {
		return s
	}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result += string(r + 32)
		} else {
			result += string(r)
		}
	}
	return result
}
