package main

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(", y = ")
	PrintNbr(points.y)
}
