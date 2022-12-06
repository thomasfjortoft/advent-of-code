package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	input, err := os.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	start := time.Now()
	fmt.Printf("Part1: %d took %s \n", part1(string(input)), time.Since(start))

	start = time.Now()
	fmt.Printf("Part2: %d took %s \n", part2(string(input)), time.Since(start))
}

func part1(input string) int {
	for i := 4; i < len(input); i++ {
		if !duplicates(input[i-3 : i+1]) {
			return i + 1
		}
	}
	return 0
}

func part2(input string) int {
	for i := 14; i < len(input); i++ {
		if !duplicates(input[i-13 : i+1]) {
			return i + 1
		}
	}
	return 0
}

func duplicates(intSlice string) bool {
	keys := make(map[rune]bool)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
		} else {
			return true
		}
	}
	return false
}
