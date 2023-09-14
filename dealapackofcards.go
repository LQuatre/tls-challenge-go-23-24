package piscine

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		PrintStr("Player " + IntToString(i+1) + ": ")
		for j := i * 3; j < i*3+3; j++ {
			PrintStr(IntToString(deck[j]))
			PrintStr(" ")
		}
		PrintStr("\n")
	}
}
