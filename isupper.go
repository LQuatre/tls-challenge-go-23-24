package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		if 'A' > r || r > 'Z' {
			return false
		}
	}
	return true
}
