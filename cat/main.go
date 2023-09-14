package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		PrintStr("ERROR: " + err.Error() + "\n")
		os.Exit(1)
	}

	for i := 0; i < len(args); i++ {
		PrintStr(args[i] + "\n")
	}

	PrintStr(string(bytes))
}
