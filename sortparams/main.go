package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Swap(a *int, b *int) {
	ValA := *a
	ValB := *b
	*a = ValB
	*b = ValA
}

func main() {
	// IN : 1 a 2 A 3 b 4 C
	// OUT :
	// 1
	// 2
	// 3
	// 4
	// A
	// C
	// a
	// b

	// sortparams/main.go

	Args := os.Args
	table := []int{}
	for i, v := range Args {
		if i > 0 {
			table = append(table, int(v[0]))
		}
	}
	n := len(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if table[j] > table[j+1] {
				Swap(&table[j], &table[j+1])
			}
		}
	}
	for _, v := range table {
		for _, c := range string(v) {
			z01.PrintRune(c)
		}
	}
}
