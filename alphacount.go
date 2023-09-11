package piscine

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' {
			count++
		}
	}
	return count
}
