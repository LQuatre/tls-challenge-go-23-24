package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	result := ""
	baseLength := 0
	for _, digit := range base {
		if digit == digit {
			baseLength++
		}
	}
	maxPower := baseLength
	if nbr < 0 {
		result = "-"
		maxPower *= -1
	}
	if baseLength > 1 {
		for nbr/maxPower >= baseLength {
			maxPower *= baseLength
		}
		for maxPower != 0 {
			result = result + string(base[nbr/maxPower])
			nbr = nbr - nbr/maxPower*maxPower
			maxPower /= baseLength
		}
		seen := map[rune]bool{}
		for _, digit := range base {
			if digit == '+' || digit == '-' {
				result = "NV"
				break
			}
			if seen[digit] == false {
				seen[digit] = true
			} else {
				result = "NV"
				break
			}
		}
	} else {
		result = "NV"
	}
	for _, digit := range result {
		z01.PrintRune(digit)
	}
}
