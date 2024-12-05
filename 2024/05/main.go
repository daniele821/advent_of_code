package main

import (
	"os"
	"strconv"
	"strings"
)

func parseFile(file string) (order map[int][]int) {
	fileByte, _ := os.ReadFile(file)
	fileStr := string(fileByte)
	phase := 0
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
				order[val0] = []int{val0}
			}
		case 1:
		}
	}
	return order
}

func part1() {
	parseFile("example1")
}

func part2() {}

func main() {
	part1()
	part2()
}
