package piscine

func IsUpper(s string) bool {
	upper := true
	for _, r := range s {
		if 'A' >= r || r >= 'Z' {
			upper = false
		}
	}
	return upper
}
