package main

import "os"

func Atoi(s string) (int, string) {
	var res int
	var sign int = 1
	for _, v := range s {
		if v == '-' && res == 0 {
			sign = -1
		} else if v == '+' && res == 0 {
			sign = 1
		} else if v < '0' || v > '9' {
			return 0, "marche pas"
		}
		res = res*10 + int(v-'0')
	}
	res *= sign
	return res, "ok"
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

	a, err := Atoi(os.Args[1])
	if err != "ok" {
		return
	}
	if a > 9223372036854775807 || a < -9223372036854775808 {
		return
	}

	b, err := Atoi(os.Args[3])
	if err != "ok" {
		return
	}
	if b > 9223372036854775807 || b < -9223372036854775808 {
		return
	}

	switch op {
	case "+":
		os.Stdout.WriteString(Itoa(a+b) + "\n")
	case "-":
		os.Stdout.WriteString(Itoa(a-b) + "\n")
	case "*":
		os.Stdout.WriteString(Itoa(a*b) + "\n")
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		os.Stdout.WriteString(Itoa(a/b) + "\n")
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		os.Stdout.WriteString(Itoa(a%b) + "\n")
	}
}
