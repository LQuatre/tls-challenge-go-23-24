package main

import "os"

func Atoi(s string) int {
	var res int
	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		res = res*10 + int(v-'0')
	}
	return res
}

func Itoa(n int) string {
	var res string
	for n > 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return res
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	op := os.Args[2]
	if op != "+" && op != "-" && op != "/" && op != "*" && op != "%" {
		return
	}

	a := Atoi(os.Args[1])
	b := Atoi(os.Args[3])

	if a == 0 && os.Args[1] != "0" {
		os.Stdout.WriteString("No Modulo by 0\n")
		return
	}
	if b == 0 && os.Args[3] != "0" {
		os.Stdout.WriteString("No Modulo by 0\n")
		return
	}

	switch op {
	case "+":
		os.Stdout.WriteString(Itoa(a + b))
	case "-":
		os.Stdout.WriteString(Itoa(a - b))
	case "*":
		os.Stdout.WriteString(Itoa(a * b))
	case "/":
		if b == 0 {
			return
		}

	case "%":
		if b == 0 {
			return
		}
		os.Stdout.WriteString(Itoa(a % b))
	}
}
