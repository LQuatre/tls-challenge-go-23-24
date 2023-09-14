package main

import (
	"fmt"
	"os"
)

func main() {
	status := true
	if len(os.Args) >= 4 {
		numChars := os.Args[2]
		files := os.Args[3:]

		for i := 0; i < len(files); i++ {
			if err := tailFile(files[i], numChars); err != nil {
				fmt.Printf("open %v: no such file or directory\n", files[i])
				status = false
				continue
			}
		}
	} else {
		status = false
	}
	if !status {
		os.Exit(1)
	}
}

func tailFile(filename string, numChars string) error {
	option := []rune(os.Args[2])
	value := IntConv(option)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println()
	fmt.Printf("==> %s <==\n", filename)

	file_stat, err := file.Stat()
	if err != nil {
		fmt.Printf("could not obtain stat for file %v\n", filename)
	}
	content := make([]byte, file_stat.Size())
	file.Read(content)
	contentinrune := []rune(string(content))
	if len(contentinrune) >= value {
		last_chars := make([]rune, value)
		for j := 0; j < len(last_chars); j++ {
			last_chars[j] = contentinrune[len(contentinrune)-value+j]
		}
		fmt.Print(string(last_chars))
	} else {
		fmt.Print(string(contentinrune))
	}
	return nil
}

func IntConv(option []rune) int {
	value := 0
	for i := 0; i < len(option); i++ {
		digit := int(option[i] - '0')
		decimal := 1
		for j := 0; j < len(option)-1-i; j++ {
			decimal *= 10
		}
		value += digit * decimal
	}
	return value
}
