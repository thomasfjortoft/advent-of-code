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

	heightMap, start, _ := getInput(input)
	
	possibleRoutes := [][3]int{}
	possibleRoutes = append(possibleRoutes, [3]int{start[0], start[1], 0})
	visited := make(map[[2]int]bool)

	for len(possibleRoutes) > 0 {
		current := possibleRoutes[0]
		possibleRoutes = possibleRoutes[1:]

		if visited[[2]int{current[0], current[1]}] {
			continue
		}

		visited[[2]int{current[0], current[1]}] = true

		if heightMap[current[0]][current[1]] == "E" {
			return current[2]
		}

		for _, v := range nextPosition {
			nextX, nextY := current[0]+v[0], current[1]+v[1]
			if nextX >= 0 && nextX < len(heightMap) && nextY >= 0 && nextY < len(heightMap[0]) {
				letterDiff := distanceBetweenLetters(heightMap[current[0]][current[1]], heightMap[nextX][nextY])

				if letterDiff <= 1 {
					possibleRoutes = append(possibleRoutes, [3]int{nextX, nextY, current[2] + 1})
				}
			}
		}
	}

	return 0
}

func part2(input string) int {

	heightMap, _, end := getInput(input)
	seen := make(map[[2]int]bool)

	possibleRoutes := [][3]int{}
	possibleRoutes = append(possibleRoutes, [3]int{end[0], end[1], 0})

	for len(possibleRoutes) > 0 {
		current := possibleRoutes[0]
		possibleRoutes = possibleRoutes[1:]

		if seen[[2]int{current[0], current[1]}] {
			continue
		}

		seen[[2]int{current[0], current[1]}] = true

		if heightMap[current[0]][current[1]] == "a" {
			return current[2]
		}

		for _, v := range nextPosition {
			nextX, nextY := current[0]+v[0], current[1]+v[1]
			if nextX >= 0 && nextX < len(heightMap) && nextY >= 0 && nextY < len(heightMap[0]) {
				letterDiff := distanceBetweenLetters(heightMap[current[0]][current[1]], heightMap[nextX][nextY])

				if letterDiff >= -1 {
					possibleRoutes = append(possibleRoutes, [3]int{nextX, nextY, current[2] + 1})
				}
			}
		}

	}

	return 0
}

var nextPosition = [4][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func distanceBetweenLetters(x, y string) int {
	if x == "S" {
		x = "a"
	}
	if y == "S" {
		y = "a"
	}
	if y == "E" {
		y = "z"
	}
	if x == "E" {
		x = "z"
	}

	r := []rune(x)
	r2 := []rune(y)

	return int(r2[0]-'a') - int(r[0]-'a')
}

func getInput(input string) ([][]string, []int, []int) {

	var hightMap [][]string
	start := make([]int, 2)
	end := make([]int, 2)
	row := strings.Split(input, "\n")

	for x := 0; x < len(row); x++ {
		colum := strings.Split(row[x], "")
		for y := 0; y < len(colum); y++ {
			if colum[y] == "S" {
				start[0] = x
				start[1] = y
			} else if colum[y] == "E" {
				end[0] = x
				end[1] = y
			}
		}
		hightMap = append(hightMap, colum)
	}

	return hightMap, start, end
}
