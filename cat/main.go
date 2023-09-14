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
		PrintStr("File name missing\n")
		return
	}
	for _, arg := range args {
		bytes, err := ioutil.ReadFile(arg)
		if err != nil {
			PrintStr("ERROR: " + err.Error() + "\n")
			continue
		}
		PrintStr(string(bytes) + "\n")
	}
}
