package piscine

import "github.com/01-edu/z01"

func AddFor(n int, prev int, result string, count *int, startNumb int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				if (len(result + string(i+48))) < startNumb {
					z01.PrintRune('0')
				}
				PrintStr(result + string(i+48))
				*count++
			} else {
				AddFor(n-1, i, result+string(i+48), count, startNumb)
			}
		}
	}
}

func PrintCombN(n int) {
	var count int = 0
	var startNumb int = n
	for i := 0; i < 10; i++ {
		if n > 1 {
			AddFor(n-1, i, string(i), &count, startNumb)
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
