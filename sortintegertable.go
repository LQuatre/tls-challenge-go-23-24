package piscine

func SortIntegerTable(table []int) {
	var newtable = table
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if newtable[i] > newtable[j] {
				Swap(&newtable[i], &newtable[j])
			}
		}
	}
}
