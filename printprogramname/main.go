package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, r := range arguments[0] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
