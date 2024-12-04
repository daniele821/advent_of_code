package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBoard(inputFile string) [][]string {
	file_byte, _ := os.ReadFile(inputFile)
	file := string(file_byte)
	board := [][]string{}
	for _, line := range strings.Split(file, "\n") {
		if line == "" {
			continue
		}
		row := []string{}
		for _, rune := range line {
			row = append(row, string(rune))
		}
		board = append(board, row)
	}
	return board
}

func countXmas(board [][]string, a, b int) int {
	size := len(board)
	count := 0
	for _, dirA := range []int{-1, 0, 1} {
		for _, dirB := range []int{-1, 0, 1} {
			if dirA == 0 && dirB == 0 {
				continue
			}
			word := ""
			for iter := range 4 {
				thisA := a + iter*dirA
				thisB := b + iter*dirB
				if thisA >= 0 && thisB >= 0 && thisA < size && thisB < size {
					word += board[thisA][thisB]
				} else {
					break
				}
			}
			if word == "XMAS" {
				count += 1
			}
		}
	}
	return count
}

func countXmas2(board [][]string, a, b int) int {
	word1 := board[a-1][b-1] + board[a][b] + board[a+1][b+1]
	word2 := board[a+1][b-1] + board[a][b] + board[a-1][b+1]
	if (word1 == "SAM" || word1 == "MAS") && (word2 == "SAM" || word2 == "MAS") {
		return 1
	}
	return 0
}

func part1() {
	board := loadBoard("input")
	acc := 0
	for a, row := range board {
		for b, char := range row {
			if char == "X" {
				acc += countXmas(board, a, b)
			}
		}
	}
	fmt.Println(acc)
}

func part2() {
	board := loadBoard("input")
	size := len(board)
	acc := 0
	for a, row := range board {
		for b, char := range row {
			if a == 0 || b == 0 || a == size-1 || b == size-1 {
				continue
			}
			if char == "A" {
				acc += countXmas2(board, a, b)
			}
		}
	}
	fmt.Println(acc)
}

func main() {
	part1()
	part2()
}
