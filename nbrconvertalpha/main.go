package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	upper := false
	for _, v := range Args {
		if v == "--upper" {
			upper = true
		}
	}
	if len(Args) == 0 || len(Args) == 1 && Args[0] == "--upper" {
		return
	}
	for _, arg := range Args {
		if arg == "--upper" {
			continue
		}
		numv := 0
		for _, v := range arg {
			numv = numv*10 + int(v-'0')
		}
		if numv >= 1 && numv <= 26 {
			if upper == false {
				z01.PrintRune(rune(numv + 96))
			} else {
				z01.PrintRune(rune(numv + 64))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
