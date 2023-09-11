package piscine

func Index(s string, toFind string) int {
	count := 0
	count2 := 0
	n := len(toFind)
	for _, r := range s {
		for _, v := range toFind {
			if r == v {
				count2++
			}
		}
		if n == count2 {
			return count - n + 1
		}
		count++
	}
	return -1
}
