package main

import (
	"fmt"
	"os"
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
	var totalScore int
	for _, l := range strings.Split(input, "\n") {
		rucksackOne := l[:len(l)/2]
		rucksackTwo := l[len(l)/2:]
		for _, letter := range rucksackOne {
			if strings.ContainsRune(rucksackTwo, letter) {
				totalScore += priority(letter)
				break
			}
		}
	}
	return totalScore
}

func part2(input string) int {
	var totalScore int
	rucksacks := make([]string, 0)
	for _, rucksack := range strings.Split(input, "\n") {
		rucksacks = append(rucksacks, rucksack)
		if len(rucksacks) == 3 {
			for _, r := range rucksacks[0] {
				if strings.ContainsRune(rucksacks[1], r) && strings.ContainsRune(rucksacks[2], r) {
					totalScore += priority(r)
					break
				}
			}
			rucksacks = make([]string, 0)
		}
	}
	return totalScore
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	return 0
}
