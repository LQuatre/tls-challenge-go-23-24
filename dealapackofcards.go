package piscine

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		PrintStr("Player ")
		PrintNbr(i + 1)
		PrintStr(": ")
		for j := i * 3; j < i*3+3; j++ {
			PrintNbr(deck[j])
			PrintStr(" ")
		}
		PrintStr("\n")
	}
}
