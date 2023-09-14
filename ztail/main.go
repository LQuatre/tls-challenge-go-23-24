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
			if err := tailFile(files[i], numChars, i); err != nil {
				fmt.Fprintln(os.Stderr, err)
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

func tailFile(filename string, numChars string, pos int) error {
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

	if pos > 3 {
		fmt.Print("\n")
	}
	fmt.Printf("==> %s <==\n", filename)

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
	return num, nil
}
