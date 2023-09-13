package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	rangeSlice := make([]int, max-min) // Créer un slice de taille max-min, un slice est un tableau dynamique
	for i := 0; i < max-min; i++ {
		rangeSlice[i] = min + i // On ajoute min à i pour avoir les valeurs de min à max
	}
	return rangeSlice
}
