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
	grid := getGridInput(input)
	total := 0

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			current := grid[row][col]

			isVisibleTop, isVisibleBottom, isVisibleRight, isVisibleLeft := true, true, true, true

			for top := row - 1; top >= 0; top-- {
				if current <= grid[top][col] {
					isVisibleTop = false
					break
				}
			}

			for bottom := row + 1; bottom < len(grid); bottom++ {
				if current <= grid[bottom][col] {
					isVisibleBottom = false
					break
				}
			}

			for left := col - 1; left >= 0; left-- {
				if current <= grid[row][left] {
					isVisibleLeft = false
					break
				}
			}

			for right := col + 1; right < len(grid); right++ {
				if current <= grid[row][right] {
					isVisibleRight = false
					break
				}
			}

			if isVisibleTop || isVisibleBottom || isVisibleLeft || isVisibleRight {
				total += 1
			}
		}
	}

	outerRing := (len(grid) * 2) + (len(grid[0]) * 2) - 4

	return total + outerRing
}

func part2(input string) int {
	grid := getGridInput(input)
	max := 0

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			current := grid[row][col]

			t, b, l, r := 0, 0, 0, 0

			for top := row - 1; top >= 0; top-- {
				t += 1
				if current <= grid[top][col] {
					break
				}
			}

			for bottom := row + 1; bottom < len(grid); bottom++ {
				b += 1
				if current <= grid[bottom][col] {
					break
				}
			}

			for left := col - 1; left >= 0; left-- {
				l += 1
				if current <= grid[row][left] {
					break
				}
			}

			for right := col + 1; right < len(grid); right++ {
				r += 1
				if current <= grid[row][right] {
					break
				}
			}

			if temp := t * b * l * r; max < temp {
				max = temp
			}

		}
	}

	return max
}

func getGridInput(input string) [][]int {
	var grid [][]int

	for _, v := range strings.Split(input, "\n") {
		var tmp []int
		for _, x := range strings.Split(v, "") {
			i, _ := strconv.Atoi(x)
			tmp = append(tmp, i)
		}
		grid = append(grid, tmp)

	}
	return grid
}
