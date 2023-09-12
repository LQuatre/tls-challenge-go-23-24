package main

import (
	"os"
)

func BasicAtoi2(s string) int {
	arrayStr := []rune(s)
	n := len(s)
	num := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return 0
		} else {
			num *= 10
			num += int(arrayStr[i]) - '0'
		}
	}
	return num
}

func main() {
	Args := os.Args
	table := []int{}
	for i := 1; i < len(Args); i++ {
		table = append(table, BasicAtoi2(Args[i]))
	}
	// flags "--upper"
	upper := false
	if Args[1] == "--upper" {
		upper = true
	}
	if upper == false {
		for i := 0; i < len(table); i++ {
			table[i] = table[i] + 96
		}
	} else {
		for i := 0; i < len(table); i++ {
			table[i] = table[i] + 64
		}
	}
	for i := 0; i < len(table); i++ {
		print(string(table[i]))
	}
	println()
}
