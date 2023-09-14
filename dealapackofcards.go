package piscine

func DealAPackOfCards(deck []int) {
	result := ""
	for i := 0; i < 4; i++ {
		result += "Player "
		result += IntToString(i + 1)
		result += ":"
		for j := i * 3; j < i*3+3; j++ {
			result += " "
			result += IntToString(deck[j])
			result += ","
		}
		result += "\n"
	}
	PrintStr(result)
}
