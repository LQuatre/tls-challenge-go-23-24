package piscine

func LoafOfBread(str string) string {
	result := ""
	count := 0
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	for i := 0; i < len(str); i++ {
		if count == 5 {
			if i != len(str)-1 {
				result += " "
			}
			count = 0
		} else {
			if str[i] == ' ' {
				continue
			}
			result += string(str[i])
			count++
		}
	}
	return result + "\n"
}
