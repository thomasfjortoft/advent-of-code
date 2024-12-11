package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type pos struct {
	row, col int
}

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
	mapData := parseInput(input)
	rows, cols := len(mapData), len(mapData[0])
	antennas := make(map[rune][]pos)
	antinodes := make(map[pos]bool)

	for r, row := range mapData {
		for c, value := range row {
			if value != '.' {
				antennas[value] = append(antennas[value], pos{r, c})
			}
		}
	}

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				p1, p2 := positions[i], positions[j]
				antiNode1 := pos{p1.row - (p2.row - p1.row), p1.col - (p2.col - p1.col)}
				antiNode2 := pos{p2.row - (p1.row - p2.row), p2.col - (p1.col - p2.col)}

				if isValid(antiNode1, rows, cols) {
					antinodes[antiNode1] = true
				}
				if isValid(antiNode2, rows, cols) {
					antinodes[antiNode2] = true
				}
			}
		}
	}

	return len(antinodes)
}

func part2(input string) int {
	mapData := parseInput(input)
	rows, cols := len(mapData), len(mapData[0])
	antennas := make(map[rune][]pos)
	antinodes := make(map[pos]bool)

	for r, row := range mapData {
		for c, value := range row {
			if value != '.' {
				antennas[value] = append(antennas[value], pos{r, c})
			}
		}
	}

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				p1, p2 := positions[i], positions[j]
				antiNode1 := pos{p1.row - (p2.row - p1.row), p1.col - (p2.col - p1.col)}
				antiNode2 := pos{p2.row - (p1.row - p2.row), p2.col - (p1.col - p2.col)}

				if isValid(antiNode1, rows, cols) {
					antinodes[antiNode1] = true
				}
				if isValid(antiNode2, rows, cols) {
					antinodes[antiNode2] = true
				}
			}
		}
	}

	return len(antinodes)
}

func isValid(p pos, rows, cols int) bool {
	return p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols
}

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mapData := make([][]rune, len(lines))
	for i, line := range lines {
		mapData[i] = []rune(line)
	}
	return mapData
}
