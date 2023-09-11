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

func CharIsMultiple(str string, r rune) bool {
	count := 0
	for _, v := range str {
		if v == r {
			count++
		}
	}
	if count > 1 {
		return true
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

func BaseIsUnique(base string) bool {
	for i, r := range base {
		for j := i + 1; j < len(base); j++ {
			if r == rune(base[j]) {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if BaseIsValid(base) {
		if nbr < 0 {
			z01.PrintRune('-')
			nbr = -nbr
		}
		if nbr >= len(base) {
			PrintNbrBase(nbr/len(base), base)
		}
		z01.PrintRune(rune(base[nbr%len(base)]))
	} else if CharIsNotUnique(base) {
		PrintStr("NV")
		return
	} else if CharIsMultiple(base, '+') || CharIsMultiple(base, '-') {
		PrintStr("NV")
	} else if BaseIsUnique(base) {
		PrintStr("NV")
	} else {
		PrintStr("NV")
		return
	}
}
