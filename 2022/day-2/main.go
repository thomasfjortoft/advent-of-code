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
		round := strings.Split(l, " ")
		totalScore += calculateScore(round[0], round[1])
	}
	return totalScore
}

func part2(input string) int {
	var totalScore int
	for _, l := range strings.Split(input, "\n") {
		round := strings.Split(l, " ")
		
		if round[0] == "A" && round[1] == "Z" {
			totalScore += calculateScore(round[0], "Y")
		} else if round[0] == "A" && round[1] == "X" {
			totalScore += calculateScore(round[0], "Z")
		} else if round[0] == "A" && round[1] == "Y" {
			totalScore += calculateScore(round[0], "X")
		} else if round[0] == "B" && round[1] == "Z" {
			totalScore += calculateScore(round[0], "Z")
		} else if round[0] == "B" && round[1] == "X" {
			totalScore += calculateScore(round[0], "X")
		} else if round[0] == "B" && round[1] == "Y" {
			totalScore += calculateScore(round[0], "Y")
		} else if round[0] == "C" && round[1] == "Z" {
			totalScore += calculateScore(round[0], "X")
		} else if round[0] == "C" && round[1] == "X" {
			totalScore += calculateScore(round[0], "Y")
		} else if round[0] == "C" && round[1] == "Y" {
			totalScore += calculateScore(round[0], "Z")
		}

	}
	return totalScore
}

func calculateScore(a string, b string) int {
	var roundScore int
	if a == "A" && b == "X" {
		roundScore += 3
	} else if a == "A" && b == "Y" {
		roundScore += 6
	} else if a == "B" && b == "Y" {
		roundScore += 3
	} else if a == "B" && b == "Z" {
		roundScore += 6
	} else if a == "C" && b == "X" {
		roundScore += 6
	} else if a == "C" && b == "Z" {
		roundScore += 3
	}

	if b == "Z" {
		roundScore += 3
	} else if b == "X" {
		roundScore += 1
	} else if b == "Y" {
		roundScore += 2

	}
	return roundScore
}
