package main

import "github.com/01-edu/z01"

type point struct {
	x rune
	y rune
}

func setPoint(ptr *point) {
	ptr.x = 42 + 48
	ptr.y = 21 + 48
}

func PrintStr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	z01.PrintRune(points.x)
	PrintStr(", y = ")
	z01.PrintRune(points.y)
	PrintStr("")
}
