package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", part1(string(input)))
	fmt.Println("Part2: ", part2(string(input)))
}

func part1(input string) int {
	var temp int
	var highest int
	for _, l := range strings.Split(input, "\n") {
		val, err := strconv.Atoi(l)
		if err != nil {
			if temp > highest {
				highest = temp
			}
			temp = 0
		}
		temp += val
	}
	return highest
}

func part2(input string) int {
	var temp int
	var elf_carrying_calories []int
	for _, l := range strings.Split(input, "\n") {
		if len(l) == 0 {
			elf_carrying_calories = append(elf_carrying_calories, temp)
			temp = 0
		} else {
			val, _ := strconv.Atoi(l)
			temp += val
		}
	}
	elf_carrying_calories = append(elf_carrying_calories, temp)
	sort.Ints(elf_carrying_calories)
	length_list := len(elf_carrying_calories)
	return elf_carrying_calories[length_list-1] + elf_carrying_calories[length_list-2] + elf_carrying_calories[length_list-3]
}
