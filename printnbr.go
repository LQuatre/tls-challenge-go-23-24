package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}
	for n != 0 {
		z01.PrintRune(rune(n%10 + 48))
		n /= 10
	}
}
