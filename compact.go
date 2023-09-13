package piscine

func Compact(ptr *[]string) int {
	var result []string
	result = *ptr
	for i := 0; i < len(result); i++ {
		if result[i] == "" {
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}
	*ptr = result
	return len(result)
}
