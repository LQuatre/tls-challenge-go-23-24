package main

import "github.com/01-edu/z01"

func main() {
	for letter := 97; letter < 123; letter++ {
		z01.PrintRune(rune(letter))
	}

	invalidRune := rune(-1)
	err := z01.PrintRune(invalidRune)
	if err == nil {
		panic("z01.PrintRune should fail with an invalid rune")
	}
	// Output:
	// 01
	// ♥
}
