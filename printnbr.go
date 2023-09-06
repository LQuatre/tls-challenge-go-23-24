package piscine

import (
	"github.com/01-edu/z01"
)

func IntToString(n int) string {
	if n == 0 { // si n est zéro, on retourne "0"
		return "0"
	}
	var s string // on crée une chaîne vide
	var neg bool // on crée un booléen pour indiquer si n est négatif
	if n < 0 {   // si n est négatif, on le rend positif et on met neg à vrai
		n = -n
		neg = true
	}
	for n > 0 { // tant que n est supérieur à zéro, on répète le processus
		s = string('0'+n%10) + s // on ajoute le reste de la division par 10 au début de la chaîne
		n = n / 10               // on divise n par 10
	}
	if neg { // si neg est vrai, on ajoute le signe moins au début de la chaîne
		s = "-" + s
	}
	return s // on retourne la chaîne finale
}

func PrintNbr(n int) {
	s := IntToString(n)   // convertit n en chaîne de caractères
	for _, r := range s { // parcourt chaque caractère de s
		z01.PrintRune(r) // affiche le caractère comme une rune
	}
}
