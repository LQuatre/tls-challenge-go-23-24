package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[0]
	for _, v := range Args[2:] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
