package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	for _, v := range Args {
		for _, c := range v {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
