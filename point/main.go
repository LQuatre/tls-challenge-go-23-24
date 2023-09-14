package main

import "github.com/01-edu/z01"

type point struct {
	x rune
	y rune
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
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
	z01.PrintRune(rune(points.x) + 48)
	PrintStr(", y = ")
	z01.PrintRune(rune(points.y) + 48)
	PrintStr("")
}
