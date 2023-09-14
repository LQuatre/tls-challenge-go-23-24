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
	for i := 0; i < len(args); i++ {
		bytes, err := ioutil.ReadFile(args[i])
		if err != nil {
			PrintStr("ERROR: open " + args[i] + ": no such file or directory\n")
			os.Exit(1)
			return
		}
		PrintStr(string(bytes))
	}
}
