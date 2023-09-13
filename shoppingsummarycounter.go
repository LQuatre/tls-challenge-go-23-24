package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	a := SplitWhiteSpaces2(str)
	for _, v := range a {
		result[v]++
	}
	return result
}
