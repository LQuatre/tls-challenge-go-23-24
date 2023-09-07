package piscine

func Abort(a, b, c, d, e int) int {
	var args = []int{a, b, c, d, e}
	var len = 5
	var temp int
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if args[j] > args[j+1] {
				temp = args[j]
				args[j] = args[j+1]
				args[j+1] = temp
			}
		}
	}
	return args[2]
}
