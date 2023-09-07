package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for _, r := range string(n) {
		z01.PrintRune(r)
	}
}
