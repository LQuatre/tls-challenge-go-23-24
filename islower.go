package piscine

func IsLower(s string) bool {
	for _, r := range s {
		if 'a' > r || r > 'z' {
			return false
		}
	}
	return true
}
