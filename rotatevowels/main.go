package main

import (
	"os"

	"github.com/01-edu/z01"
)

// rotateVowels is a function that rotates vowels in a string

func Swap(vow []rune) {
	for i := 0; i < len(vow)/2; i++ {
		vow[i], vow[len(vow)-i-1] = vow[len(vow)-i-1], vow[i]
	}
}

func main() {
	argruments := os.Args[1:]
	length := 0

	for i := range argruments {
		length = i + 1
	}

	if length != 0 {
		str := ""
		first := false

		for _, arg := range argruments {
			if first {
				str += " "
			}
			first = true
			str += arg
		}

		runes := []rune(str)
		var pos []int
		var vow []rune

		for i, r := range runes {
			if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
				pos = append(pos, i)
				vow = append(vow, r)
			}
		}
		Swap(vow)

		for i, v := range pos {
			runes[v] = vow[i]
		}

		for _, r := range runes {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
