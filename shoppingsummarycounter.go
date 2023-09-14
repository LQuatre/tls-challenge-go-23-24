package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	wordList := SplitWhiteSpaces(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}
