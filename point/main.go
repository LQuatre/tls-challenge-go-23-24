package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
}

func PrintRune(r int) {
	if r >= 10 {
		PrintRune(r / 10)
	}
	res := rune(r%10 + 48)
	z01.PrintRune(res)
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintRune(points.x)
	PrintStr(", y = ")
	PrintRune(points.y)
}