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

type pos struct{ x, y int }

var _map [][]rune

func part1(input string) int {
	startPos := pos{0, 0}

	// Parse input and initialize map
	lines := strings.Split(input, "\n")
	_map = make([][]rune, len(lines))
	for rowIndex, line := range lines {
		_map[rowIndex] = []rune(line)
		for colIndex, char := range line {
			if char == '^' {
				startPos = pos{colIndex, rowIndex}
			}
		}
	}

	_, steps := walker(startPos, _map)
	return steps
}

func part2(input string) int {
	var _map [][]rune
	startPos := pos{0, 0}

	// Parse input and initialize map
	lines := strings.Split(input, "\n")
	_map = make([][]rune, len(lines))
	for rowIndex, line := range lines {
		_map[rowIndex] = []rune(line)
		for colIndex, char := range line {
			if char == '^' {
				startPos = pos{colIndex, rowIndex}
			}
		}
	}
	numberOfLoops := 0
	for y, row := range _map {
		for x, c := range row {
			if c != '^' && c != '#' {
				_map[y][x] = '#'
				var isLoop, _ = walker(startPos, _map)
				if isLoop { numberOfLoops++ }
				_map[y][x] = '.'
			}
		}
	}

	return numberOfLoops
}

func walker(startPos pos, _map [][]rune) (bool, int) {
	directions := []pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	walked := 0
	visited := make(map[pos]bool)
	currentPos := startPos
	visited[currentPos] = true
	for {
		nextPos := pos{currentPos.x + directions[0].x, currentPos.y + directions[0].y}
		if !visited[currentPos] {
			visited[currentPos] = true
		}

		if currentPos.x == 0 || currentPos.y == 0 || currentPos.x == len(_map[0])-1 || currentPos.y == len(_map)-1 {
			return false, len(visited)
		}
		if walked-len(visited) > 1000 {
			return true, len(visited)
		}

		if _map[nextPos.y][nextPos.x] == '#' {
			directions = append(directions[1:], directions[0])
		} else {
			currentPos = nextPos
			walked++
		}
	}
}
