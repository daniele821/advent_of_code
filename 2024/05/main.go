package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(file string) (map[int][]int, map[int][]int, [][]int) {
	fileByte, _ := os.ReadFile(file)
	fileStr := string(fileByte)
	phase := 0
	order1 := map[int][]int{}
	order2 := map[int][]int{}
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
			if val, ok := order1[val0]; ok {
				order1[val0] = append(val, val1)
			} else {
				order1[val0] = []int{val1}
			}
			if val, ok := order2[val1]; ok {
				order2[val1] = append(val, val0)
			} else {
				order2[val1] = []int{val0}
			}
		case 1:
		}
	}
	return order1, order2, nil
}

func part1() {
	order1, order2, lists := parseFile("example1")
	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(lists)
}

func part2() {}

func main() {
	part1()
	part2()
}
