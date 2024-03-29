package piscine

import (
	"github.com/01-edu/z01"
)

func UintToString(n uint) string {
	var res string
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	return res
}

func IntToString(n int) string {
	if n < 0 {
		z01.PrintRune(45)
		return UintToString(uint(-n))
	}
	return UintToString(uint(n))
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}
	for n != 0 {
		z01.PrintRune(rune(n%10 + 48))
		n /= 10
	}
}
