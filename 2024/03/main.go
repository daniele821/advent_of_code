package main

import (
	"fmt"
	"os"
)

func part1() {
	file_byte, _ := os.ReadFile("example")
	file := string(file_byte)
	fmt.Println(file)
}

func part2() {
}

func main() {
	part1()
	part2()
}
