package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Parse command line arguments
	count := flag.Int("c", 0, "number of bytes to print from the end of the file")
	flag.Parse()

	// Get the list of files to process
	files := flag.Args()

	// Loop over each file and print the last count bytes
	for _, file := range files {
		// Print the file name if there are multiple files
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", file)
		}

		// Open the file
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v\n", file, err)
			continue
		}
		defer f.Close()

		// Get the file size
		fi, err := f.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "stat %s: %v\n", file, err)
			continue
		}
		size := fi.Size()

		// Read the last count bytes
		offset := size - int64(*count)
		if offset < 0 {
			offset = 0
		}
		_, err = f.Seek(offset, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "seek %s: %v\n", file, err)
			continue
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", file, err)
			continue
		}

		// Print the data
		fmt.Printf("%s", data)

		// Print a newline between files
		if len(files) > 1 {
			fmt.Println()
		}
	}
}
