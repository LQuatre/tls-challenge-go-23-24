package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var s string = string(n)
	for _, r := range s {
		z01.PrintRune(r)
	}
}
