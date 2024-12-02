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
	var safeCount int
	for _, line := range strings.Split(input, "\n") {
		levels := strings.Fields(line)
		if len(levels) == 0 {
			continue
		}
		var elements []int
		for _, level := range levels {
			element, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			elements = append(elements, element)
		}
		if isSafe(elements) {
			safeCount++
		}
	}
	return safeCount
}

func part2(input string) int {
	var safeCount int
	for _, line := range strings.Split(input, "\n") {
		levels := strings.Fields(line)
		if len(levels) == 0 {
			continue
		}
		var elements []int
		for _, level := range levels {
			element, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			elements = append(elements, element)
		}
		if isSafe2(elements) {
			safeCount++
		}

	}
	return safeCount
}

func isSafe2(elements []int) bool {
	if isSafe(elements) {
		return true
	}
	for i := 0; i < len(elements); i++ {
		e := remove(elements, i)
		if isSafe(e) {
			return true
		}
	}
	return false
}

func remove(elements []int, i int) []int {
	l := len(elements)
	if i == 0 {
		return elements[1:]
	}
	if i == l-1 {
		return elements[0 : l-1]
	}
	e := make([]int, len(elements))
	copy(e, elements)
	return append(e[0:i], e[i+1:]...)
}

func isSafe(elements []int) bool {
	isSafe := true
	for i := 1; i < len(elements); i++ {
		prev := elements[i-1]
		curr := elements[i]
		diff := curr - prev
		if diff < -3 || diff > 3 || diff == 0 {
			isSafe = false
			break
		}
		if i == 1 {
			continue
		}
		prevDiff := prev - elements[i-2]
		if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
			isSafe = false
			break
		}
	}
	return isSafe
}
