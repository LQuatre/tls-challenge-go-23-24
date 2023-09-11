package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, r := range strs {
		if i == 0 {
			result += r
		} else {
			result += sep + r
		}
	}
	return result
}
