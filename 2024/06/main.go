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

func printGrid(grid [][]string) {
	for _, line := range grid {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
}
func countGrid(grid [][]string) {
	count := 0
	for _, line := range grid {
		for _, char := range line {
			if char == "X" {
				count++
			}
		}
	}
	fmt.Println(count)
}

func solve(file string) {
	fileStr, _ := os.ReadFile(file)
	lines := strings.Split(string(fileStr), "\n")
	square := grid(lines)
	tracked := grid(lines)

	size := len(square)

	x, y := 0, 0
	for x1, elem := range square {
		for y1, elem2 := range elem {
			if elem2 == "^" {
				x, y = x1, y1
			}
		}
	}

	dirx, diry := -1, 0

	for x >= 0 && y >= 0 && x < size && y < size {
		tracked[x][y] = "X"
		x += dirx
		y += diry
		if !(x >= 0 && y >= 0 && x < size && y < size) {
			break
		}
		if square[x][y] == "#" {
			x -= dirx
			y -= diry
			if dirx == -1 && diry == 0 {
				dirx = 0
				diry = 1
			} else if dirx == 0 && diry == 1 {
				dirx = 1
				diry = 0
			} else if dirx == 1 && diry == 0 {
				dirx = 0
				diry = -1
			} else if dirx == 0 && diry == -1 {
				dirx = -1
				diry = 0
			}
		}
	}
	countGrid(tracked)

}

func main() {
	solve("example")
	solve("input")
}
