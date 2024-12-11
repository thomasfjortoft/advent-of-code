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
	var scores int
	for r, row := range values {
		for c, value := range row {
			if value != 0 {
				continue
			}
			var visited = make(map[pos]bool)
			var queue []pos
			queue = append(queue, pos{r, c})
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				if values[p.row][p.col] == 9 {
					if !visited[p] { scores++ }
					visited[p] = true
				}
				for _, move := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					if p.row + move.row >= 0 && p.row + move.row < len(values) &&
					   p.col + move.col >= 0 && p.col + move.col < len(values[p.row])	&&
					   (values[p.row + move.row][p.col + move.col] - values[p.row][p.col]) == 1 {
						queue = append(queue, pos{p.row + move.row, p.col + move.col})
					}
				}
			}
		}
	}
	return scores
}

func part2(input string) int {
	values := parseInput(input)
	var scores int
	for r, row := range values {
		for c, value := range row {
			if value != 0 {
				continue
			}
			var visited = make(map[pos]bool)
			var queue []pos
			queue = append(queue, pos{r, c})
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				if values[p.row][p.col] == 9 {
					scores++
					visited[p] = true
				}
				for _, move := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					if p.row + move.row >= 0 && p.row + move.row < len(values) &&
					   p.col + move.col >= 0 && p.col + move.col < len(values[p.row])	&&
					   (values[p.row + move.row][p.col + move.col] - values[p.row][p.col]) == 1 {
						queue = append(queue, pos{p.row + move.row, p.col + move.col})
					}
				}
			}
		}
	}
	return scores
}

func parseInput(input string) [][]int {
	rows := strings.Split(input, "\n")
	output := make([][]int, len(rows))
	for i := range output {
		output[i] = make([]int, len(rows[0]))
	}
	for i, row := range rows {
		values := strings.Split(row, "")
		for j, value := range values {
			s, _ := strconv.Atoi(value)
			output[i][j] = s
		}
	}
	return output
}

type pos struct { row, col int }