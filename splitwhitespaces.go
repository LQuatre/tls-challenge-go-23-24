package piscine

func AddTemp(result *[]string, temp *string) {
	if len(*temp) > 0 {
		*result = append(*result, *temp)
		*temp = ""
	}
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var temp string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			AddTemp(&result, &temp)
			continue
		}
		temp += string(s[i])
		if i == len(s)-1 {
			AddTemp(&result, &temp)
		}
	}
	return result
}
