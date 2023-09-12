package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[0]
	for k, v := range Args {
		if k > 1 {
			z01.PrintRune(rune(v))
		}
	}
}
