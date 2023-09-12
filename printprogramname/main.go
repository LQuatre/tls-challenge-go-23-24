package main

import (
	"os"

	"github.com/01-edu/z01"
)

func DernierSlashEtLaSuite(s string) string {
	var index int
	for i, r := range s {
		if r == '\\' {
			index = i
		}
	}
	return s[index:]
}

func RemoveExe(s string) string {
	var index int
	for i, r := range s {
		if r == '.' {
			index = i
		}
	}
	return s[:index]
}

func main() {
	Args := os.Args[0]
	Args = DernierSlashEtLaSuite(Args)
	Args = RemoveExe(Args)
	for _, r := range Args {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
