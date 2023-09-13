package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d : ", i+1)
		for j := i * 3; j < (i+1)*3; j++ {
			fmt.Printf("%d ", deck[j])
		}
		fmt.Println()
	}
}
