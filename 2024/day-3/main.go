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
	var total int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		total += getMulArgs(line)
	}
	return total
}

func part2(input string) int {
	var total int
	lines := strings.Split(input, "\n")
	inputLines := strings.Join(lines, "")
	ignoreParts := strings.Split(inputLines, "don't()")
	total += getMulArgs(ignoreParts[0])
	for _, ignorePart := range ignoreParts[0:] {
		endDont := strings.Index(ignorePart, "do()")
		if endDont != -1 {
			ignoreArgs := ignorePart[endDont:]
			total += getMulArgs(ignoreArgs)
		}
	}

	return total
}

func getMulArgs(input string) int {
	total := 0
	parts := strings.Split(input, "mul(")
	for _, part := range parts[1:] {
		endIdx := strings.Index(part, ")")
		if endIdx != -1 {
			mulArgs := part[:endIdx]
			nums := strings.Split(mulArgs, ",")
			if len(nums) == 2 {
				x, err1 := strconv.Atoi(nums[0])
				y, err2 := strconv.Atoi(nums[1])
				if err1 == nil && err2 == nil {
					total += x * y
				}
			}
		}
	}
	return total
}
