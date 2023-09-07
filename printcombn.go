package piscine

import (
	"github.com/01-edu/z01"
)

func UintToString(n uint) string {
	var res string
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	return res
}

func IntToString(n int) string {
	if n < 0 {
		z01.PrintRune('-')
		return UintToString(uint(-n))
	}
	return UintToString(uint(n))
}

func AddFor(n int, prev int, result string, count *int) {
	for i := 0; i <= 9; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				for _, r := range result {
					z01.PrintRune(r)
				}
				*count++
			} else {
				AddFor(n-1, i, result+string(i+48), count)
			}
		}
	}
}

func PrintCombN(n int) {
	var count int = 0
	for i := 0; i <= 9; i++ {
		if n > 1 {
			AddFor(n-1, i, IntToString(i), &count)
		} else {
			if count > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(i) + 48)
			count++
		}
	}
	z01.PrintRune('\n')
}
