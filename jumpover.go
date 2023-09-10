package piscine

import "github.com/01-edu/z01"

var first = false

func JumpOver(str string) string {
	if first == false {
		first = true
	} else {
		z01.PrintRune(10)
	}
	var result string
	// jump two str characters and add the third to result
	for i := 2; i < StrLen(str); i += 3 {
		result += string(str[i])
	}
	return result
}
