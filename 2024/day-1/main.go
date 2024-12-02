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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(input string) int {
	var left_number []int
	var right_number []int
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, "   ")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		left_number = append(left_number, left)
		right_number = append(right_number, right)
	}
	sort.Ints(left_number)
	sort.Ints(right_number)

	var total int
	for i := range left_number {
		diff := left_number[i] - right_number[i]
		total += abs(diff)
	}

	return total
}

func part2(input string) int {
	var left_number []int
	rightCount := make(map[int]int)
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, "   ")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		left_number = append(left_number, left)
		rightCount[right]++
	}

	var total int
	for i := range left_number {
		diff := left_number[i] * rightCount[left_number[i]]
		total += abs(diff)
	}

	return total
}
