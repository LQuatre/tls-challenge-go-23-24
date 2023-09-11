package piscine

func isValidBase(base string) bool {
	baseLength := len(base)
	if baseLength < 2 {
		return false
	}
	if Index(base, "-") != -1 || Index(base, "+") != -1 {
		return false
	}
	for i, char := range base {
		if Index(base, string(char)) != i {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}
	baseLength := len(base)
	result := 0
	for _, digit := range s {
		for i, char := range base {
			if digit == char {
				result = result*baseLength + i
				break
			}
		}
	}
	return result
}
