package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	len := 0
	for range a {
		len++
	}
	if len == 0 || len == 1 {
		return true
	}
	if f(a[0], a[1]) < 0 {
		for i := 1; i < len-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	} else if f(a[0], a[1]) > 0 {
		for i := 1; i < len-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	}
	return true
}
