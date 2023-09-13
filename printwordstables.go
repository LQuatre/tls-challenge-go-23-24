package piscine

func PrintWordsTables(a []string) {
	result := ""
	temp := ""
	for i := 0; i < len(a); i++ {
		temp = a[i]
		result = temp + "\n"
	}
	PrintStr(result)
}
