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
	total := 0
	ops := []string{"+", "*"}
	for _, puzzle := range parseInput(input) {
		//testNumber := puzzle.testValue
		if generateEquations(puzzle.testValue, puzzle.values, 0, 0, ops) {
			total += puzzle.testValue
		}

	}
	return total
}

func part2(input string) int {
	total := 0
	ops := []string{"+", "*", "|"}
	for _, puzzle := range parseInput(input) {
		//testNumber := puzzle.testValue
		if generateEquations(puzzle.testValue, puzzle.values, 0, 0, ops) {
			total += puzzle.testValue
		}

	}
	return total
}

type puzzle struct {
	testValue int
	values    []int
}

func parseInput(input string) []puzzle {
	var puzzles []puzzle = make([]puzzle, 0)
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		parts := strings.Split(row, ": ")
		testValue, _ := strconv.Atoi(parts[0])
		numbers := strings.Split(parts[1], " ")
		var values []int
		for _, num := range numbers {
			value, _ := strconv.Atoi(num)
			values = append(values, value)
		}
		puzzles = append(puzzles, puzzle{
			testValue: testValue,
			values:    values,
		})
	}
	return puzzles
}

func generateEquations(testNumber int, numbers []int, index int, current int, ops []string) bool {
	if index == len(numbers) {
		return current == testNumber
	}

	for _, op := range ops {
		next := numbers[index]
		var newCurrent int
		if op == "+" {
			newCurrent = current + next
		} else if op == "*" {
			newCurrent = current * next
		} else if op == "|" {
			newCurrent, _ = strconv.Atoi(fmt.Sprintf("%d%d", current, next))
		}
		if generateEquations(testNumber, numbers, index+1, newCurrent, ops) {
			return true
		}
	}
	return false
}
