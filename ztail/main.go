package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . -c <num_chars> <file1> [<file2> ...]")
		os.Exit(1)
	}

	numChars := os.Args[2]
	files := os.Args[3:]

	for _, file := range files {
		fmt.Printf("==> %s <==\n", file)

		if err := tailFile(file, numChars); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if len(files) > 1 {
			fmt.Println()
		}
	}
}

func tailFile(filename string, numChars string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()

	numBytesToRead, err := parseNumChars(numChars)
	if err != nil {
		return err
	}

	if fileSize < numBytesToRead {
		numBytesToRead = fileSize
	}

	_, err = file.Seek(fileSize-numBytesToRead, 0)
	if err != nil {
		return err
	}

	buf := make([]byte, numBytesToRead)
	_, err = file.Read(buf)
	if err != nil {
		return err
	}

	fmt.Printf("%s", buf)
	return nil
}

func parseNumChars(numChars string) (int64, error) {
	var num int64
	for _, char := range numChars {
		if char >= '0' && char <= '9' {
			num = num*10 + int64(char-'0')
		} else {
			return 0, fmt.Errorf("invalid number of characters: %s", numChars)
		}
	}
	return num + 2, nil
}
