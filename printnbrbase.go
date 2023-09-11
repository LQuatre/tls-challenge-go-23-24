package piscine

import "github.com/01-edu/z01"

func checkBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] || base[i] == '+' || base[i] == '-' {
				return false
			}
		}
	}
	return true
}

func IsUnique(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] || str[i] == '+' || str[i] == '-' {
				return false
			}
		}
	}
	return true
}

func ContainsNotValidChar(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == '+' || str[i] == '-' {
			return true
		}
	}
	return false
}

func PrintNbrBase(nbr int, base string) {
	if !checkBase(base) || !IsUnique(base) || ContainsNotValidChar(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	if nbr >= len(base) {
		PrintNbrBase(nbr/len(base), base)
	}
	z01.PrintRune(rune(base[nbr%len(base)]))
}
