package piscine

import (
	"github.com/01-edu/z01"
)



func AddFor(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				// print(string(result))
				if len(result) <= 1 {
					z01.PrintRune('0')
				}
				for _, r := range result {
					z01.PrintRune(rune(r))
				}
				z01.PrintRune(rune(i) + 48)
				*count++
			} else {
				AddFor(n-1, i, result+string(i+48), count)
			}
		}
	}
}

func PrintCombN(n int) {
	var count int = 0
	for i := 0; i < 10; i++ {
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
