package main

import (
	"fmt"
	"os"
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
	return 0
}

func part2(input string) int {
	return 0
}