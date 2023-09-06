package main

import "github.com/01-edu/z01"

func UintToString(n uint) string {
	var res string
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	return res
}

func IntToString(n int) string {
	if n < 0 {
		z01.PrintRune('-')
		return UintToString(uint(-n))
	}
	return UintToString(uint(n))
}

func PrintCombN(n int) {
	if n > 0 || n < 10 {

	}
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
