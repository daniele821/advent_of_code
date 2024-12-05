package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseFile(file string) (order1, order2 map[int][]int, precedences [][]int, updates [][]int) {
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

func removeElement(elem int, slice []int) []int {
	index := slices.Index(slice, elem)
	if index == -1 {
		panic("unable to remove elem")
	}
	return append(slice[:index], slice[index+1:]...)
}

func incorrectOrder(order1, order2 map[int][]int, precedences [][]int, line []int) int {
	lineCopy := make([]int, len(line))
	copy(lineCopy, line)
	newVals := []int{}

	for len(lineCopy) > 0 {
		for i, value := range lineCopy {
			afterVal := order2[value]
			count := 0
			for i2, value2 := range lineCopy {
				if i2 == i {
					continue
				}
				if slices.Contains(afterVal, value2) {
					count++
				}
			}
			if count == 0 {
				newVals = append(newVals, value)
				lineCopy = removeElement(value, lineCopy)
				break
			}
		}
	}
	fmt.Println(newVals)
	return newVals[len(newVals)/2]
}

func part1() {
	_, _, precedences, lists := parseFile("input")
	acc := 0
	for _, list := range lists {
		tmp, errs := correctOrder(precedences, list)
		if errs == 0 {
			acc += tmp
		}
	}
	fmt.Println(acc)
}

func part2() {
	order1, order2, precedences, lists := parseFile("input")
	acc := 0
	for _, list := range lists {
		_, errs := correctOrder(precedences, list)
		if errs != 0 {
			acc += incorrectOrder(order1, order2, precedences, list)
		}
	}
	fmt.Println(acc)
}

func main() {
	// part1()
	part2()
}
