package main

import (
	"fmt"
	"os"
	"strconv"
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
	values := parseInput(input)
	positionsMap := make(map[string][]pos)
	var visited = make(map[pos]bool)
	for r, row := range values {
		for c, value := range row {
			if _, ok := visited[pos{r, c}]; ok {
				continue
			}
			var queue []pos
			queue = append(queue, pos{r, c})
			mapKey := string(value) + strconv.Itoa(r) + strconv.Itoa(c)
			var singleVisited = make(map[pos]bool)

			for len(queue) > 0 {
				currentPos := queue[0]
				queue = queue[1:]
				singleVisited[currentPos] = true
				positionsMap[mapKey] = append(positionsMap[mapKey], currentPos)
				for _, move := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					nextPos := pos{currentPos.row + move.row, currentPos.col + move.col}
					if nextPos.row >= 0 && nextPos.row < len(values) &&
						nextPos.col >= 0 && nextPos.col < len(values[currentPos.row]) {
						if (values[nextPos.row][nextPos.col] == values[currentPos.row][currentPos.col] && !singleVisited[pos{nextPos.row, nextPos.col}] && !visited[nextPos]) {
							queue = append(queue, pos{nextPos.row, nextPos.col})
							visited[nextPos] = true
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, positions := range positionsMap {
		perimeter := 0
		for _, p := range positions {
			for _, move := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				neighbor := pos{p.row + move.row, p.col + move.col}
				if !contains(positions, neighbor) {
					perimeter++
				}
			}
		}
		sum += perimeter * len(positions)
	}
	return sum
}

func part2(input string) int {

	return 0
}

func parseInput(input string) [][]rune {
	rows := strings.Split(input, "\n")
	output := make([][]rune, len(rows))
	for i := range output {
		output[i] = make([]rune, len(rows[0]))
	}
	for i, row := range rows {
		values := strings.Split(row, "")
		for j, value := range values {
			output[i][j] = []rune(value)[0]
		}
	}
	return output
}

type pos struct{ row, col int }

func contains(slice []pos, item pos) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}