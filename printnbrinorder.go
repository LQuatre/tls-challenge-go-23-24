package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(s int) {
	if s < 0 {
		return
	}
	if s == 0 {
		z01.PrintRune('0')
		return
	}
	var arr []int
	for s > 0 {
		arr = append(arr, s%10)
		s /= 10
	}
	for i := range arr {
		for j := range arr {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	for _, v := range arr {
		z01.PrintRune(rune(v + 48))
	}
}
