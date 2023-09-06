package main

import "github.com/01-edu/z01"

func main() {
	for letter := 97; letter < 123; letter++ {
		z01.PrintRune(rune(letter))
	}
}
