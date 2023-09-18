package main

import "os"

func Atoi(s string) (int, string) {
	// gestion du signe
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	if s[0] == '+' {
		s = s[1:]
	}
	res := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')
		} else {
			return 0, "error"
		}
	}
	return res * sign, "ok"
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	s := ""
	for n > 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	return sign + s
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
