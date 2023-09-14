package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	for _, arg := range args {
		fmt.Println(arg)
	}
}
