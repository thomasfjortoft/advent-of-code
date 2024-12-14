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
	fmt.Printf("Part1: %d took %s \n", part1(string(input), 101, 103), time.Since(start))

	start = time.Now()
	fmt.Printf("Part2: %d took %s \n", part2(string(input)), time.Since(start))
}

func part1(input string, width, hight int) int {
	RUN_TIME := 100
	puzzles := parseInput(input)
	for i := 0; i < RUN_TIME; i++ {
		moveRobots(puzzles, width, hight)
	}

	return calculateSafety(puzzles, width, hight)
}
func part2(input string) int {
	width := 101
	hight := 103
	RUN_TIME := 7339
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	puzzles := parseInput(input)
	for i := 1; i <= RUN_TIME; i++ {
		moveRobots(puzzles, width, hight)
		file.WriteString(fmt.Sprintf("Step: %d", i))
		for i := 0; i < width; i++ {
			for j := 0; j < hight; j++ {
				if puzzleAt(puzzles, i, j) {
					file.WriteString("#")
				} else {
					file.WriteString(".")
				}
			}
			file.WriteString("\n")
		}

	}

	return 0
}

func puzzleAt(puzzles []puzzle, x, y int) bool {
	for _, robot := range puzzles {
		if robot.px == x && robot.py == y {
			return true
		}
	}
	return false
}

func calculateSafety(puzzles []puzzle, width, hight int) int {

	quadrantUpRight := 0
	quadrantUpLeft := 0
	quadrantDownRight := 0
	quadrantDownLeft := 0
	midCol := width / 2
	midRow := hight / 2

	for _, robot := range puzzles {
		if robot.px > midCol && robot.py < midRow {
			quadrantUpRight++
		} else if robot.px < midCol && robot.py < midRow {
			quadrantUpLeft++
		} else if robot.px > midCol && robot.py > midRow {
			quadrantDownRight++
		} else if robot.px < midCol && robot.py > midRow {
			quadrantDownLeft++
		}
	}
	return quadrantDownLeft * quadrantDownRight * quadrantUpLeft * quadrantUpRight
}

func moveRobots(puzzles []puzzle, width, hight int) {
	for i := range puzzles {
		robot := &puzzles[i]
		newX := robot.px + robot.vx
		newY := robot.py + robot.vy
		if newX < 0 {
			newX = width + newX
		}
		if newX >= width {
			newX = newX - width
		}
		if newY < 0 {
			newY = hight + newY
		}
		if newY >= hight {
			newY = newY - hight
		}

		robot.px = newX
		robot.py = newY
	}
}

func parseInput(input string) []puzzle {
	var puzzles []puzzle
	for _, robot := range strings.Split(input, "\n") {
		var p puzzle
		fmt.Sscanf(robot, "p=%d,%d v=%d,%d", &p.px, &p.py, &p.vx, &p.vy)
		puzzles = append(puzzles, p)
	}
	return puzzles
}

type puzzle struct {
	px, py, vx, vy int
}
