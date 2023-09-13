package piscine

func AddTemp2(result *string, temp *string) {
	if len(*temp) > 0 {
		if len(*result) > 0 {
			*result += " "
		}
		*result += *temp
		*temp = ""
	}
}

func Split(s, sep string) string {
	var result string
	var temp string
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			AddTemp2(&result, &temp)
			i += len(sep) - 1
			continue
		}
		temp += string(s[i])
		if i == len(s)-1 {
			AddTemp2(&result, &temp)
		}
	}
	return result
}
