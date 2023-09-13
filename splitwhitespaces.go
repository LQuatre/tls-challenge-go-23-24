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

func SplitWhiteSpaces2(s string) []string {
	var result []string
	var temp string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			for j := i + 1; j < len(s); j++ {
				if s[j] == ' ' || s[j] == '\n' || s[j] == '\t' {
					temp += string(s[j])
					AddTemp(&result, &temp)
					i++
				} else {
					break
				}
			}
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
