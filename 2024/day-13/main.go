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
	fmt.Printf("Part1: %d took %s \n", part1(string(input)), time.Since(start))

	start = time.Now()
	fmt.Printf("Part2: %d took %s \n", part2(string(input)), time.Since(start))
}

func part1(input string) int {
	total := 0
	for _, p := range parseInput(input) {
		value := findValues(p.ax, p.ay, p.bx, p.by, p.px, p.py, 0)
		total += int(value)
	}

	return total
}

func part2(input string) int {
	total := 0
	for _, p := range parseInput(input) {
		value := findValues(p.ax, p.ay, p.bx, p.by, p.px, p.py, 10000000000000)
		total += int(value)
	}

	return total
}

func parseInput(input string) []puzzle {
	var puzzles []puzzle
	for _, block := range strings.Split(input, "\n\n") {
		var p puzzle
		fmt.Sscanf(block, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &p.ax, &p.ay, &p.bx, &p.by, &p.px, &p.py)
		puzzles = append(puzzles, p)
	}
	return puzzles
}

func findValues(ax, ay, bx, by, px, py, offSett int) (uint64) {
	b := float64(ax*(py+offSett) - ay*(px+offSett)) / float64(ax*by - ay*bx)
	a := (float64(px+offSett) - float64(bx)*float64(b)) / float64(ax)
	if a == float64(uint64(a)) && b == float64(uint64(b)) {
		return 3*uint64(a) + uint64(b)
	}
	return 0
}

type puzzle struct {
	ax, ay, bx, by, px, py int
}
