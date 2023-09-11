package piscine

import "github.com/01-edu/z01"

func Count(str, substr string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == substr[0] {
			if str[i:i+len(substr)] == substr {
				count++
			}
		}
	}
	return count
}

func PrintNbrBase(nbr int, str string) {
	indx := 0
	for _, res := range str {
		if string(res) == "-" || string(res) == "+" || Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2 {
		PrintStr("NV")
	} else if -9223372036854775808 <= nbr && nbr >= 9223372036854775807 {
		PrintStr("NV")
	} else {
		if nbr < 0 {
			z01.PrintRune('-')
			nbr = -nbr
		}
		if nbr >= len(str) {
			PrintNbrBase(nbr/len(str), str)
		}
		z01.PrintRune(rune(str[nbr%len(str)]))
	}
}
