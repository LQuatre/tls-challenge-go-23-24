package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune(rune(84))
		z01.PrintRune('\n')
		return
	}
	z01.PrintRune(rune(70))
	z01.PrintRune('\n')
}
