package main

import (
	"fmt"
	"os"
)

// $ go run . --insert=4321 --order asdad
// 1234aadds
// $ go run . --insert=4321 asdad
// asdad4321
// $ go run . asdad
// asdad
// $ go run . --order 43a21
// 1234a
// $ go run .
// --insert
//   -i
//          This flag inserts the string into the string passed as argument.
// --order
//   -o
//          This flag will behave like a boolean, if it is called it will order the argument.
// $ go run . -h
// --insert
//   -i
//          This flag inserts the string into the string passed as argument.
// --order
//   -o
//          This flag will behave like a boolean, if it is called it will order the argument.
// $ go run . --help
// --insert
//   -i
//          This flag inserts the string into the string passed as argument.
// --order
//   -o
//          This flag will behave like a boolean, if it is called it will order the argument.

func Help() {
	fmt.Println("--insert\n")
	fmt.Println("  -i\n")
	fmt.Println("	   This flag inserts the string into the string passed as argument.\n")
	fmt.Println("--order\n")
	fmt.Println("  -o\n")
	fmt.Println("	   This flag will behave like a boolean, if it is called it will order the argument	.`\n")
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
			return
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
