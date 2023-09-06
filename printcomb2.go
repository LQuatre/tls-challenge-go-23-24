package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i != j {
						if (i != j || i != l) && (k != l) {
							// XX-X					--XX
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(' ')
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
							if i != '9' || j != '9' || k != '9' || l != '9' {
								z01.PrintRune(rune(','))
								z01.PrintRune(rune(' '))
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
