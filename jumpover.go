package piscine

func JumpOver(str string) string {
	var result string
	for i := 2; i < StrLen(str); i += 3 {
		result += string(str[i])
	}
	return result + "\n"
}
