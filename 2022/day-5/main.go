package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	input, err := os.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	start := time.Now()
	fmt.Printf("Part1: %s took %s \n", part1(string(input)), time.Since(start))

	start = time.Now()
	fmt.Printf("Part2: %s took %s \n", part2(string(input)), time.Since(start))
}

func part1(input string) string {
	s := strings.Split(input, "\n\n")
	stacks, steps := readMap(s[0]), s[1]
	for _, l := range strings.Split(steps, "\n") {
		var count, from, to int
		fmt.Sscanf(l, "move %d from %d to %d", &count, &from, &to)
		for i := 0; i < count; i++ {
			n := len(stacks[from-1]) - 1
			topElement := stacks[from-1][n]
			stacks[from-1] = stacks[from-1][:n]
			stacks[to-1] = append(stacks[to-1], topElement)
		}
	}
	return getTop(stacks)
}

func part2(input string) string {
	s := strings.Split(input, "\n\n")
	stacks, steps := readMap(s[0]), s[1]
	for _, l := range strings.Split(steps, "\n") {
		var count, from, to int
		fmt.Sscanf(l, "move %d from %d to %d", &count, &from, &to)
		var temp []string
		for i := 0; i < count; i++ {
			n := len(stacks[from-1]) - 1
			topElement := stacks[from-1][n]
			stacks[from-1] = stacks[from-1][:n]
			temp = append(temp, topElement)
		}

		for i := len(temp); i > 0; i-- {
			stacks[to-1] = append(stacks[to-1], temp[i-1])
		}
	}
	return getTop(stacks)
}

func readMap(input string) [][]string {
	lines := strings.Split(input, "\n")
	testArray := strings.Fields(lines[len(lines)-1])
	stacks := make([][]string, len(testArray))
	for i := len(lines) - 2; i >= 0; i-- {
		t := strings.Split(lines[i], "")
		for row := 1; row < len(lines[i]); row += 4 {
			if t[row] != " " {
				n := (row - 1) / 4
				stacks[n] = append(stacks[n], t[row])
			}
		}
	}
	return stacks
}

func getTop(stacks [][]string) string {
	message := ""
	for i := range stacks {
		n := len(stacks[i]) - 1
		message += stacks[i][n]
	}
	return message
}
