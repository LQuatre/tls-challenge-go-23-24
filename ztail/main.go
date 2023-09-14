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

	fmt.Println()
	fmt.Printf("==> %s <==\n", filename)

	file_stat, err := file.Stat()
	if err != nil {
		fmt.Printf("could not obtain stat for file %v\n", filename)
	}
	content := make([]byte, file_stat.Size())
	file.Read(content)
	contentinrune := []rune(string(content))
	if len(contentinrune) >= int(numBytesToRead) {
		last_chars := make([]rune, numBytesToRead)
		for j := 0; j < len(last_chars); j++ {
			last_chars[j] = contentinrune[len(contentinrune)-int(numBytesToRead)+j]
		}
		fmt.Print(string(last_chars))
	} else {
		fmt.Print(string(contentinrune))
	}
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
