package piscine

import (
	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	var digits []int

	if n == 0 {
		return "0"
	}

	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	var res string

	for i := len(digits) - 1; i >= 0; i-- {
		res += string(digits[i] + 48)
	}

	return res
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	var digits []int

	s := Itoa(n)

	for _, m := range s {
		digits = append(digits, int(m-'0'))
	}

	for _, val := range digits {
		z01.PrintRune(rune(val + 48))
	}
}
