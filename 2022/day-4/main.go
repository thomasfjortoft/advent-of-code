package main

import (
	"fmt"
	"os"
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
	var totalScore int
	for _, l := range strings.Split(input, "\n") {
		assignments := strings.Split(l, ",")
		pairOne := strings.Split(assignments[0], "-")
		pairTwo := strings.Split(assignments[1], "-")
		p11, _ := strconv.Atoi(pairOne[0])
		p12, _ := strconv.Atoi(pairOne[1])
		p21, _ := strconv.Atoi(pairTwo[0])
		p22, _ := strconv.Atoi(pairTwo[1])

		// 6-6,4-6
		if p11 >= p21 && p11 <= p22 && p12 <= p22 && p12 >= p21 {
			totalScore += 1
		} else if p21 >= p11 && p21 <= p12 && p22 <= p12 && p22 >= p11 {
			totalScore += 1
		}
	}
	return totalScore
}

func part2(input string) int {
	var totalScore int
	for _, l := range strings.Split(input, "\n") {
		assignments := strings.Split(l, ",")
		pairOne := strings.Split(assignments[0], "-")
		pairTwo := strings.Split(assignments[1], "-")
		p11, _ := strconv.Atoi(pairOne[0])
		p12, _ := strconv.Atoi(pairOne[1])
		p21, _ := strconv.Atoi(pairTwo[0])
		p22, _ := strconv.Atoi(pairTwo[1])

		// 5-7,7-9
		if (p11 >= p21 && p11 <= p22) || (p12 <= p22 && p12 >= p21) {
			totalScore += 1
		} else if (p21 >= p11 && p21 <= p12) || (p22 <= p12 && p22 >= p11) {
			totalScore += 1
		}
	}
	return totalScore
}
