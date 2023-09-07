package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	args := []int{a, b, c, d, e}
	sort.Sort(sort.IntSlice(args))
	return args[2]
}
