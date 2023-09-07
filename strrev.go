package piscine

func StrRev(s string) string {
	var result string
	for _, r := range s {
		result += string(r)
	}
	return result
}
