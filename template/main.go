package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(file string) {
	fileStr, _ := os.ReadFile(file)
	lines := strings.Split(string(fileStr), "\n")

	// write code here
	fmt.Println(lines)
}

func main() {
	solve("example")
	solve("input")
}
