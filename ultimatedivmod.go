package piscine

func UltimateDivMod(a *int, b *int) {
	var c int
	var d int
	c = *a
	d = *b
	*a = c / d
	*b = c % d
}
