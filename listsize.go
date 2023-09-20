package piscine

func ListSize(l *List) int {
	count := 0
	for l != nil {
		count++
	}
	return count
}
