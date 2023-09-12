package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args
	var result string
	for i, v := range Args {
		if i > 0 {
			result = v + "\n" + result
		}
	}
	for _, c := range result {
		z01.PrintRune(c)
	}
}
