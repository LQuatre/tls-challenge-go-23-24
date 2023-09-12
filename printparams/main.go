package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args
	for i, v := range Args {
		if i > 0 {
			for _, c := range v {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
	}
}
