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
	var count int
	for i := range os.Args {
		count = i
	}
	if count == 0 {
		str, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			PrintStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		} else {
			PrintStr(string(str))
		}
	} else {
		if len(os.Args) == 0 {
			return
		}
		os.Args = os.Args[1:]
		for i := range os.Args {
			str, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				PrintStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			} else {
				PrintStr(string(str))
			}
		}
	}
}
