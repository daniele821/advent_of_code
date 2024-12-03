package main

import (
	"fmt"
	"os"
	"strconv"
)

func isVal(str string) bool {
	if len(str) != 1 {
		panic("invalid input: str len must be 1!")
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

func part1() {
	fileByte, _ := os.ReadFile("input")
	file := string(fileByte)
	length := len(file)
	acc := 0
	for index, rune := range file {
		rune_str := string(rune)
		if rune_str == "m" {
			area := file[index:min(length, index+12)]
			if area[:4] != "mul(" {
				continue
			}
			area = area[4:]
			val1, l, found := startsWithNum(area)
			if !found {
				continue
			} else {
				area = area[l:]
			}
			if area[:1] != "," {
				continue
			}
			area = area[1:]
			val2, l, found := startsWithNum(area)
			if !found {
				continue
			} else {
				area = area[l:]
			}
			if area[:1] != ")" {
				continue
			}
			acc += val1 * val2
		}
	}
	fmt.Println(acc)
}

func part2() {
}

func main() {
	part1()
	part2()
}
