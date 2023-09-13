package piscine

import (
	"math"
)

func CreateDigitMap(base string) map[rune]int {
	digitMap := make(map[rune]int)
	for i, digit := range base {
		digitMap[digit] = i
	}
	return digitMap
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digitMap := CreateDigitMap(digits)

	baseFromInt := len(baseFrom)
	nbrInt := 0
	for i, digit := range nbr {
		nbrInt += digitMap[digit] * int(math.Pow(float64(baseFromInt), float64(len(nbr)-i-1)))
	}

	baseToInt := len(baseTo)
	result := ""
	for nbrInt > 0 {
		result = string(digits[nbrInt%baseToInt]) + result
		nbrInt /= baseToInt
	}

	if result == "" {
		result = "0"
	}

	return result
}
