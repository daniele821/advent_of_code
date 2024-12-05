package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(file string) map[int][]int {
	fileByte, _ := os.ReadFile(file)
	fileStr := string(fileByte)
	phase := 0
	order := map[int][]int{}
	for _, line := range strings.Split(fileStr, "\n") {
		if line == "" {
			phase++
			continue
		}
		switch phase {
		case 0:
			vals := strings.Split(line, "|")
			if len(vals) != 2 {
				panic("invalid argoment")
			}
			val0, _ := strconv.Atoi(vals[0])
			val1, _ := strconv.Atoi(vals[1])
			if val, ok := order[val0]; ok {
				order[val0] = append(val, val1)
			} else {
				order[val0] = []int{val1}
			}
		case 1:
		}
	}
	return order
}

func part1() {
	order := parseFile("example1")
	fmt.Println(order)
}

func part2() {}

func main() {
	part1()
	part2()
}
