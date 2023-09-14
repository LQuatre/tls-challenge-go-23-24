package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	wordList := SplitWhiteSpaces(str)
	result := make(map[string]int)
	for _, word := range wordList {
		_, ok := result[word]
		if ok {
			result[word] += 1
		} else {
			result[word] = 1
		}
	}
	return result
}
