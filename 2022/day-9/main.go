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
	visited := make(map[string]int)
	visited["x0y0"] = 1
	headX, headY, tailX, tailY := 0, 0, 0, 0
	for _, command := range strings.Split(input, "\n") {
		var direction string
		var steps int
		fmt.Sscanf(command, "%s %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				headX += 1
				if isInReach(headX, headY, tailX, tailY) {
					continue
				}
				tailX, tailY = headX-1, headY
			case "L":
				headX -= 1
				if isInReach(headX, headY, tailX, tailY) {
					continue
				}
				tailX, tailY = headX+1, headY
			case "U":
				headY += 1
				if isInReach(headX, headY, tailX, tailY) {
					continue
				}
				tailX, tailY = headX, headY-1
			case "D":
				headY -= 1
				if isInReach(headX, headY, tailX, tailY) {
					continue
				}
				tailX, tailY = headX, headY+1
			}

			key := fmt.Sprintf("x%dy%d", tailX, tailY)
			_, ok := visited[key]
			if !ok {
				visited[key] = 1
				continue
			}
			visited[key]++
		}

	}

	return len(visited)
}

func isInReach(headX, headY, tailX, tailY int) bool {
	xDiff := headX - tailX
	yDiff := headY - tailY

	return xDiff > -2 && xDiff < 2 && yDiff > -2 && yDiff < 2
}

func part2(input string) int {
	return 0
}
