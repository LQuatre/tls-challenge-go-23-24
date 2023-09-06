package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func PrintNbr(n int) { // Les commentaire vont décrire ce que fait la fonction
	var digits []int // On déclare un tableau de int

	s := strconv.Itoa(n) // On convertit le int en string

	for _, m := range s { // On parcourt le string
		digits = append(digits, int(m-'0')) // On ajoute chaque chiffre dans le tableau
	}

	for _, val := range digits { // On parcourt le tableau
		ascii := rune(val + 48) // On convertit le int en rune
		z01.PrintRune(ascii)    // On affiche le rune
	}
}
