package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: ztail -c num file1 file2 ...")
		os.Exit(1)
	}

	c := os.Args[1]
	if c != "-c" {
		fmt.Fprintln(os.Stderr, "usage: ztail -c num file1 file2 ...")
		os.Exit(1)
	}

	num := os.Args[2]
	if num[0] == '-' {
		fmt.Fprintln(os.Stderr, "usage: ztail -c num file1 file2 ...")
		os.Exit(1)
	}

	for i := 3; i < len(os.Args); i++ {
		if i > 3 {
			fmt.Println()
		}
		fmt.Printf("==> %s <==\n", os.Args[i])
		f, err := os.Open(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v\n", os.Args[i], err)
			continue
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "stat %s: %v\n", os.Args[i], err)
			continue
		}

		if stat.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is a directory\n", os.Args[i])
			continue
		}

		if numInt, err := strconv.Atoi(num); err == nil {
			if stat.Size() < int64(numInt) {
				numInt = int(stat.Size())
			}
			if _, err := f.Seek(-int64(numInt), io.SeekEnd); err != nil {
				fmt.Fprintf(os.Stderr, "seek %s: %v\n", os.Args[i], err)
				continue
			}
		}

		if _, err := io.Copy(os.Stdout, f); err != nil {
			fmt.Fprintf(os.Stderr, "copy %s: %v\n", os.Args[i], err)
			continue
		}
	}
}
