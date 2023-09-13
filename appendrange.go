package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var result []int
	for i := min; i < max; i++ {
		result = append(result, i) // append() permet d'ajouter un élément à un slice, un slice est un tableau dynamique
	}
	return result
}
