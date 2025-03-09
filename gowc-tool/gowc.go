package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines, words, bytes := 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		bytes += len(line) + 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "readomg standard output:", err)
	}

	fmt.Printf("%d lines %d words %d bytes", lines, words, bytes)
}
