package piscine

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		PrintStr("Player: ")
		for j := 0; j < len(deck); j++ {
			PrintStr(string(deck[j]))
			PrintStr(" ")
		}
		PrintStr("\n")
	}
}
