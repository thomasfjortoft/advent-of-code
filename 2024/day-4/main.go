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

func part1(input string) int {
	var horizontal int
	var vertical int
	var diagonal int
	lines := strings.Split(input, "\n")
	for rowIndex, line := range lines {
		for colIndex, char := range line {
			// Check for horizontal XMAS or SAMX
			if strings.HasPrefix(line[colIndex:], "XMAS") || strings.HasPrefix(line[colIndex:], "SAMX") {
				horizontal++
			}
			// Check for vertical XMAS or SAMX
			if rowIndex < len(lines)-3 {
				if string(char) == "X" && string(lines[rowIndex+1][colIndex]) == "M" && string(lines[rowIndex+2][colIndex]) == "A" && string(lines[rowIndex+3][colIndex]) == "S" {
					vertical++
				}
				if string(char) == "S" && string(lines[rowIndex+1][colIndex]) == "A" && string(lines[rowIndex+2][colIndex]) == "M" && string(lines[rowIndex+3][colIndex]) == "X" {
					vertical++
				}
			}
			// Check for diagonal XMAS or SAMX
			if rowIndex < len(lines)-3 && colIndex < len(line)-3 {
				if string(char) == "X" && string(lines[rowIndex+1][colIndex+1]) == "M" && string(lines[rowIndex+2][colIndex+2]) == "A" && string(lines[rowIndex+3][colIndex+3]) == "S" {
					diagonal++
				}
				if string(char) == "S" && string(lines[rowIndex+1][colIndex+1]) == "A" && string(lines[rowIndex+2][colIndex+2]) == "M" && string(lines[rowIndex+3][colIndex+3]) == "X" {
					diagonal++
				}
			}
			// Check for anti-diagonal XMAS or SAMX
			if rowIndex < len(lines)-3 && colIndex > 2 {
				if string(char) == "X" && string(lines[rowIndex+1][colIndex-1]) == "M" && string(lines[rowIndex+2][colIndex-2]) == "A" && string(lines[rowIndex+3][colIndex-3]) == "S" {
					diagonal++
				}
				if string(char) == "S" && string(lines[rowIndex+1][colIndex-1]) == "A" && string(lines[rowIndex+2][colIndex-2]) == "M" && string(lines[rowIndex+3][colIndex-3]) == "X" {
					diagonal++
				}
			}
		}
	}
	println(horizontal, vertical, diagonal)
	return horizontal + vertical + diagonal
}

func part2(input string) int {
	total := 0
	lines := strings.Split(input, "\n")
	for rowIndex, line := range lines {
		for colIndex, char := range line {
			if rowIndex < len(line)-1 && colIndex < len(line)-1 && rowIndex > 0 && colIndex > 0 {
				if string(char) == "A" {
					if string(lines[rowIndex-1][colIndex-1]) == "M" && string(lines[rowIndex-1][colIndex+1]) == "S" && string(lines[rowIndex+1][colIndex-1]) == "M" && string(lines[rowIndex+1][colIndex+1]) == "S" {
						total++
					}
					if string(lines[rowIndex-1][colIndex-1]) == "S" && string(lines[rowIndex-1][colIndex+1]) == "M" && string(lines[rowIndex+1][colIndex-1]) == "S" && string(lines[rowIndex+1][colIndex+1]) == "M" {
						total++
					}
					if string(lines[rowIndex-1][colIndex-1]) == "M" && string(lines[rowIndex-1][colIndex+1]) == "M" && string(lines[rowIndex+1][colIndex-1]) == "S" && string(lines[rowIndex+1][colIndex+1]) == "S" {
						total++
					}
					if string(lines[rowIndex-1][colIndex-1]) == "S" && string(lines[rowIndex-1][colIndex+1]) == "S" && string(lines[rowIndex+1][colIndex-1]) == "M" && string(lines[rowIndex+1][colIndex+1]) == "M" {
						total++
					}
					
				}
			}
		}
	}
	return total
}
