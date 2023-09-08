package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if table[j] > table[j+1] {
				Swap(&table[j], &table[j+1])
			}
		}
	}
}
