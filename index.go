package piscine

func Index(s string, toFind string) int {
	lenS := len(s)
	lenToFind := len(toFind)
	if lenToFind == 0 || lenS == 0 || lenToFind > lenS {
		return 0
	}
	for i := 0; i < lenS; i++ {
		if s[i] == toFind[0] {
			if i+lenToFind > lenS {
				return -1
			}
			if toFind == s[i:i+lenToFind] {
				return i
			}
		}
	}
	return -1
}
