package main

import (
	"fmt"
	"os"
	"strconv"
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
	total := 0
	rules := strings.Split(input, "\n\n")[0]
	pages := strings.Split(input, "\n\n")[1]
	for _, page := range strings.Split(pages, "\n") {
		pageCorrect := true
		pageNumbers := strings.Split(page, ",")
		for pageNumberIndex, pageNumber := range pageNumbers {
			if pageNumberIndex == len(pageNumbers)-1 {
				break
			}
			if !strings.Contains(rules, pageNumber+"|"+string(pageNumbers[pageNumberIndex+1])) {
				pageCorrect = false
			}
		}
		if pageCorrect {
			total += calculateMiddleNumberSum(pageNumbers)
		}
	}

	return total

}

func part2(input string) int {
	total := 0
	rules := strings.Split(input, "\n\n")[0]
	pages := strings.Split(input, "\n\n")[1]
	for _, page := range strings.Split(pages, "\n") {
		pageNumbers := strings.Split(page, ",")
		pageCorrect := true
		RestartLoop:
			for pageNumberIndex := 0; pageNumberIndex < len(pageNumbers)-1; pageNumberIndex++ {
				pageNumber := pageNumbers[pageNumberIndex]
				if !strings.Contains(rules, pageNumber+"|"+pageNumbers[pageNumberIndex+1]) {
					pageNumbers[pageNumberIndex], pageNumbers[pageNumberIndex+1] = pageNumbers[pageNumberIndex+1], pageNumbers[pageNumberIndex]
					pageCorrect = false
					goto RestartLoop
				}
			}
		if !pageCorrect {
			total += calculateMiddleNumberSum(pageNumbers)
		}
	}

	return total
}
func calculateMiddleNumberSum(pageNumbers []string) int {
	middleIndex := len(pageNumbers) / 2
	middleNumber := pageNumbers[middleIndex]
	middleNumberInt, err := strconv.Atoi(middleNumber)
	if err != nil {
		panic(err)
	}
	return middleNumberInt
}
