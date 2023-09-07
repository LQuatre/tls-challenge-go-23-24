package piscine

func StrRev(s string) string {
	result := ""
	l := StrLen(s)
	for i := l; i != 0; i-- {
		result = i + result
	}
	return result
}
