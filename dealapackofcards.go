package piscine

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		PrintStr("Player: ")
		for i := 0; i < len(deck); i++ {
			PrintStr(myItoa(deck[i]))
			if i != len(deck)-1 {
				PrintStr(" ")
			}
		}
		PrintStr("\n")
	}
}
