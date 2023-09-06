package main

import "github.com/01-edu/z01"

func main() {
	for letter := 122; letter > 96; letter-- {
		z01.PrintRune(rune(letter))
	}
	z01.PrintRune('\n')
}
