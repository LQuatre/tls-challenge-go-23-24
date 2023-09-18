package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, value := range a {
		result = append(result, f(value))
	}
	return result
}
