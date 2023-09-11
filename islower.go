package piscine

func IsLower(s string) bool {
	lower := true
	for _, r := range s {
		if 'a' > r || r > 'z' {
			lower = false
		}
	}
	return lower
}
