package main

import (
	"fmt"
	"os"
)

func StringToInt(s string) int {
	var result int
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		result = result*10 + int(c-'0')
	}
	return result
}

func main() {
	// Get the list of files to read
	args := os.Args[1:]
	// Loop through each args[i] and print the last 10 bytes
	for i := 2; i < len(args); i++ {
		fileContents, err := os.ReadFile(args[i])
		if len(fileContents) > 1 {
			fmt.Printf("==> %s <==\n", args[i])
		}

		// Open the args[i]
		f, err := os.Open(args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v\n", args[i], err)
			continue
		}
		defer f.Close()

		// Get the args[i] size
		fi, err := f.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "stat %s: %v\n", args[i], err)
			continue
		}
		size := fi.Size()

		// Read the last args[1] bytes of the args[i]
		var offset int64
		if size > int64(StringToInt(args[1])+2) {
			offset = size - int64(StringToInt(args[1])+2)
		}

		b := make([]byte, StringToInt(args[1])+2)
		_, err = f.ReadAt(b, offset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", args[i], err)
			continue
		}

		// Print the last 10 bytes of the args[i]
		fmt.Printf("%s", b)
		if b[len(b)-1] != '\n' {
			fmt.Println()
		}
	}
}
