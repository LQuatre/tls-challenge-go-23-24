package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var temp string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			result = append(result, temp)
			temp = ""
			continue
		}
		temp += string(s[i])
		if i == len(s)-1 {
			result = append(result, temp)
			temp = ""
		}
	}
	return result
}
