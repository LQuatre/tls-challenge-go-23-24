package main

import "github.com/01-edu/z01"

func main() {
	for number := 48; number < 58; number++ {
		z01.PrintRune(rune(number))
	}
	z01.PrintRune('\n')
}
