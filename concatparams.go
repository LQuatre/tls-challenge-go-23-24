package piscine

func ConcatParams(args []string) string {
	result := ""
	for k, v := range args {
		if k == 0 {
			result = v
		} else {
			result = result + "\n" + v
		}
	}
	return result
}
