package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isVal(str string) bool {
	if len(str) != 1 {
		panic("invalid: str len must be 1!")
	}
	runes := []rune(str)
	return runes[0] >= 48 && runes[0] <= 57
}

func startsWithNum(str string) (value, length int, found bool) {
	res := ""
	for _, rune := range str {
		runeStr := string(rune)
		if isVal(runeStr) {
			res += runeStr
		} else {
			break
		}
	}
	if len(res) == 0 || len(res) >= 4 {
		return -1, -1, false
	}
	val, _ := strconv.Atoi(res)
	return val, len(res), true
}

func Accumulate(area string) int {
	if area[:4] != "mul(" {
		return 0
	}
	area = area[4:]
	val1, l, found := startsWithNum(area)
	if !found {
		return 0
	} else {
		area = area[l:]
	}
	if area[:1] != "," {
		return 0
	}
	area = area[1:]
	val2, l, found := startsWithNum(area)
	if !found {
		return 0
	} else {
		area = area[l:]
	}
	if area[:1] != ")" {
		return 0
	}
	return val1 * val2
}

func part1() {
	fileByte, _ := os.ReadFile("input")
	file := string(fileByte)
	length := len(file)
	acc := 0
	for index, rune := range file {
		rune_str := string(rune)
		if rune_str == "m" {
			area := file[index:min(length, index+12)]
			acc += Accumulate(area)
		}
	}
	fmt.Println(acc)
}

func part2() {
	fileByte, _ := os.ReadFile("input")
	file := string(fileByte)
	length := len(file)
	acc := 0
	enabled := true
	for index, rune := range file {
		rune_str := string(rune)
		rem := file[index:]
		if strings.HasPrefix(rem, "do()") {
			fmt.Println("ENABLE: ", rem[:4])
			enabled = true
		}
		if strings.HasPrefix(rem, "don't()") {
			fmt.Println("DISABLE: ", rem[:7])
			enabled = false
		}
		if rune_str == "m" && enabled {
			area := file[index:min(length, index+12)]
			acc += Accumulate(area)
		}
	}
	fmt.Println(acc)
}

func main() {
	part1()
	part2()
}
