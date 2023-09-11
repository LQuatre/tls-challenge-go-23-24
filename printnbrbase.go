package piscine

import "github.com/01-edu/z01"

func BaseIsValid(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if BaseIsValid(base) {
		if nbr < 0 {
			PrintStr("-")
			nbr = -nbr
		}
		if nbr >= len(base) {
			PrintNbrBase(nbr/len(base), base)
		}
		z01.PrintRune(rune(base[nbr%len(base)]))
	} else {
		PrintStr("NV")
	}
}
