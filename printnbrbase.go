package piscine

import "github.com/01-edu/z01"

func CharIsNotUnique(str string) bool {
	for i, r := range str {
		for j := i + 1; j < len(str); j++ {
			if r == rune(str[j]) {
				return true
			}
		}
	}
	return false
}

func BaseIsValid(base string) bool {
	if base == "" {
		return false
	}
	for i, r := range base {
		if r == '-' || r == '+' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !BaseIsValid(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	} else {
		if nbr < 0 {
			z01.PrintRune('-')
			nbr = -nbr
		}
		if nbr >= len(base) {
			PrintNbrBase(nbr/len(base), base)
		}
		z01.PrintRune(rune(base[nbr%len(base)]))
	}
	if CharIsNotUnique(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
}
