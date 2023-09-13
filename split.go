package piscine

func Split(s, sep string) []string {
	var result []string
	var temp string
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			AddTemp(&result, &temp)
			i += len(sep) - 1
			continue
		}
		temp += string(s[i])
		if i == len(s)-1 {
			AddTemp(&result, &temp)
		}
	}
	return result
}
