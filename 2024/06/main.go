package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(file string) []string {
	lines, _ := os.ReadFile(file)
	return strings.Split(string(lines), "\n")
}

func solve(file string) {
	lines := readLines(file)

	// write code here
	fmt.Println(lines)
}

func main() {
	solve("example")
	solve("input")
}
