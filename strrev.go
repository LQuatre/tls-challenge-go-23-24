package piscine

func StrRev(s string) string {
	result := ""
	for _, r := range s {
		result = string(r) + result
	}
	return result
}
