package piscine

func CreateDigitMap(base string) map[rune]int {
	digitMap := make(map[rune]int)
	for i, digit := range base {
		digitMap[digit] = i
	}
	return digitMap
}

func Pow(a, b float64) float64 {
	if b == 0 {
		return 1
	}
	return a * Pow(a, b-1)
}

func baseFromInt(nbr int, baseFrom string, digitMap map[rune]int) string {
	result := ""
	for nbr > 0 {
		result = string(baseFrom[nbr%len(baseFrom)]) + result
		nbr /= len(baseFrom)
	}
	return result
}

func baseToInt(nbr string, baseTo string, digitMap map[rune]int) int {
	baseToInt := len(baseTo)
	nbrInt := 0
	for i, digit := range nbr {
		nbrInt += digitMap[digit] * int(Pow(float64(baseToInt), float64(len(nbr)-i-1)))
	}
	return nbrInt
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digitMap := CreateDigitMap(digits)

	nbrInt := baseToInt(nbr, baseFrom, digitMap)
	result := baseFromInt(nbrInt, baseTo, digitMap)

	if result == "" {
		result = "0"
	}

	return result
}
