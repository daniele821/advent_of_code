package main

import (
	"fmt"
	"os"
	"strings"
)

func grid(lines []string) [][]string {
	grid := [][]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		gridLine := []string{}
		for _, char := range line {
			gridLine = append(gridLine, string(char))
		}
		grid = append(grid, gridLine)
	}
	return grid
}

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
