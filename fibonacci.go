package piscine

func Fibonacci(index int) int {
	switch {
	case index == -1:
		return -1
	case index == 0:
		return 0
	case index == 1:
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
