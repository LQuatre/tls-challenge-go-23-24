package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Swap(a *[]rune, b *[]rune) {
	ValA := *a
	ValB := *b
	*a = ValB
	*b = ValA
}

func main() {
	Args := os.Args
	table := []string{}
	for i := 1; i < len(Args); i++ {
		table = append(table, Args[i])
	}
	tableRune := [][]rune{}
	for i := 0; i < len(table); i++ {
		tableRune = append(tableRune, []rune(table[i]))
	}
	n := len(tableRune)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if tableRune[j][0] > tableRune[j+1][0] {
				Swap(&tableRune[j], &tableRune[j+1])
			}
		}
	}
	for i := 0; i < len(tableRune); i++ {
		for j := 0; j < len(tableRune[i]); j++ {
			z01.PrintRune(tableRune[i][j])
		}
		z01.PrintRune('\n')
	}
}
