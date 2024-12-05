package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseFile(file string) (order1 map[int][]int, order2 map[int][]int, precedences [][]int, updates [][]int) {
	fileByte, _ := os.ReadFile(file)
	fileStr := string(fileByte)
	phase := 0
	order1 = map[int][]int{}
	order2 = map[int][]int{}
	precedences = [][]int{}
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
			precedences = append(precedences, []int{val0, val1})
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
			vals := strings.Split(line, ",")
			valsInt := []int{}
			for _, val := range vals {
				value, _ := strconv.Atoi(val)
				valsInt = append(valsInt, value)
			}
			updates = append(updates, valsInt)
		}
	}
	return order1, order2, precedences, updates
}

func correctOrder(precedences [][]int, line []int) (acc, errs int) {
	for _, prec := range precedences {
		index1 := slices.Index(line, prec[0])
		index2 := slices.Index(line, prec[1])
		if index1 >= 0 && index2 >= 0 {
			if index1 >= index2 {
				errs++
			}
		}
	}
	return line[len(line)/2], errs
}

func part1() {
	_, _, precedences, lists := parseFile("example1")
	acc := 0
	for _, list := range lists {
		tmp, errs := correctOrder(precedences, list)
		if errs == 0 {
			acc += tmp
		}
	}
	fmt.Println(acc)
}

func part2() {}

func main() {
	part1()
	part2()
}
