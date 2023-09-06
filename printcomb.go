package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i != j && i != k && j != k {
					if i < j && i < k && j < k {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						if i != '7' || j != '8' || k != '9' {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						} else {
							z01.PrintRune(rune('\n'))
						}
					}
				}
			}
		}
	}
}

func main() {
	PrintComb()
}
