package piscine

import "fmt"

const nEnd = 8

var position = [nEnd]int{}

func isInDanger(queennumber, rowposition int) bool {
	for i := 0; i < queennumber; i++ {
		other_row_pos := position[i]
		if other_row_pos == rowposition || other_row_pos == rowposition-(queennumber-i) || other_row_pos == rowposition+(queennumber-i) {
			return false
		}
	}
	return true
}

func resolverPuzzle(n int) {
	if n == nEnd {
		for i := 0; i < nEnd; i++ {
			fmt.Print(position[i] + 1)
		}
		fmt.Print("\n")
	} else {
		for i := 0; i < nEnd; i++ {
			if isInDanger(n, i) {
				position[n] = i
				resolverPuzzle(n + 1)
			}
		}
	}
}

func EightQueens() {
	resolverPuzzle(0)
}
