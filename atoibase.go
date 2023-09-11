package piscine

func isValidBase(base string) bool {
	baseLength := 0
	for _, digit := range base {
		if digit == digit {
			baseLength++
		}
	}
	if baseLength < 2 {
		return false
	}
	seen := map[rune]bool{}
	for _, digit := range base {
		if digit == '+' || digit == '-' {
			return false
		}
		if seen[digit] == false {
			seen[digit] = true
		} else {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}
	baseLength := 0
	for _, digit := range base {
		if digit == digit {
			baseLength++
		}
	}
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
