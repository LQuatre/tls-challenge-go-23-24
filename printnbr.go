package piscine

import (
	"github.com/01-edu/z01"
)

func Int64ToString(n int) string {
	var res string
	var neg bool
	if n < 0 {
		neg = true
		n = -n
	}
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	if neg {
		res = "-" + res
	}
	return res
}

func PrintNbr(n int) {
	s := Int64ToString(n)
	for _, r := range s {
		z01.PrintRune(r)
	}
}
