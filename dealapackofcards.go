package piscine

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		PrintStr("Player" + Itoa(i+1) + ": ")
		for j := i * 3; j < (i+1)*3; j++ {
			PrintStr(Itoa(deck[j]) + " ")
		}
		PrintStr("\n")
	}
}
