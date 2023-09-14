package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards (deck []int) {
	for i := 0; i < 4; i++ {
		PrintStr("Player: ")
		for j := 0; j < 13; j++ {
			z01.PrintRune(rune(deck[j]) + 48)
			PrintStr(" ")
		}
		PrintStr("\n")
	}
}
