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
	for i := 0; i < len(args); i++ {
		bytes, err := ioutil.ReadFile(args[i])
		if err != nil {
			PrintStr("ERROR: " + err.Error() + "\n")
			PrintStr("exit status " + string(i+1+'0'))
			return
		}
		PrintStr(string(bytes))
	}
}
