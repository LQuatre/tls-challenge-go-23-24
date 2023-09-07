package piscine

func StrLen(s string) int {
	var count int
	for i := range s {
		if i == i {
			count++
		}
	}
	return count
}
