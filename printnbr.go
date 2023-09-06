package piscine

import (
	"github.com/01-edu/z01"
)

func IntToString(n int) string {
	if n < 0 { // si n est négatif, on ajoute le signe moins au début
		return "-" + IntToString(-n)
	}
	if n == 0 { // si n est zéro, on retourne "0"
		return "0"
	}
	if n < 10 { // si n est inférieur à 10, on retourne le caractère correspondant
		return string('0' + n)
	}
	return IntToString(n/10) + string('0'+n%10) // sinon, on divise n par 10 et on ajoute le reste à la fin
}

func PrintNbr(n int) {
	s := IntToString(n)   // convertit n en chaîne de caractères
	for _, r := range s { // parcourt chaque caractère de s
		z01.PrintRune(r) // affiche le caractère comme une rune
	}
}
