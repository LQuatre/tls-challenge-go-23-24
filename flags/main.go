package main

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == substr[0] {
			if len(s)-i >= len(substr) {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
		}
	}
	return false
}

func Swap(k *rune, v *rune) {
	l := *k
	*k = *v
	*v = l
}

func Sort(s string) string {
	var result string
	var runes []rune
	for _, v := range s {
		runes = append(runes, v)
	}
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes)-1; j++ {
			if runes[j] > runes[j+1] {
				Swap(&runes[j], &runes[j+1])
			}
		}
	}
	for _, v := range runes {
		result += string(v)
	}
	return result
}

func main() {
	Args := os.Args[1:]
	var Text string
	var Insert string
	var Order bool
	var Added bool
	if len(Args) == 0 {
		Help()
	}
	for _, v := range Args {
		if Contains(v, "--help") || Contains(v, "-h") {
			Help()
		} else if Contains(v, "--insert=") {
			for _, v := range v[9:] {
				Insert += string(v)
			}
			Added = true
		}
		if Contains(v, "-i=") {
			for _, v := range v[3:] {
				Insert += string(v)
			}
			Added = true
		}
		if Contains(v, "--order") {
			Order = true
		}
		if Contains(v, "-o") {
			Order = true
		}
		if !Contains(v, "--") && !Contains(v, "-") {
			Text += v
			if Added {
				Text += Insert
			}
			if Order {
				Text = Sort(Text)
			}
			fmt.Println(Text)
		}
	}
}
