package piscine

func Pow(a, b float64) float64 {
	if b == 0 {
		return 1
	}
	return a * Pow(a, b-1)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	baseFromLen := len(baseFrom)
	baseToLen := len(baseTo)
	nbrLen := len(nbr)
	decimal := 0
	for i := nbrLen - 1; i >= 0; i-- {
		for j := 0; j < baseFromLen; j++ {
			if nbr[i] == baseFrom[j] {
				decimal += j * int(Pow(float64(baseFromLen), float64(nbrLen-i-1)))
			}
		}
	}
	result := ""
	for decimal > 0 {
		result = string(baseTo[decimal%baseToLen]) + result
		decimal /= baseToLen
	}
	if result == "" {
		result = string(baseTo[0])
	}
	return result
}
