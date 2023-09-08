package piscine

func Swap(a *int, b *int) {
	ValA := *a
	ValB := *b
	*a = ValB
	*b = ValA
}
