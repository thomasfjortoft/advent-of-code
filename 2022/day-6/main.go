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
	var answer int
	list := strings.Split(input, "")
	temp := list[:4]
	for i := 4; i < len(list); i++ {
		if !duplicates(temp) {
			answer = i
			break
		}
		temp = list[i-3 : i+1]
	}
	return answer
}

func part2(input string) int {
	var answer int
	list := strings.Split(input, "")
	temp := list[:14]
	for i := 14; i < len(list); i++ {
		if !duplicates(temp) {
			answer = i
			break
		}
		temp = list[i-13 : i+1]
	}
	return answer
}

func duplicates(intSlice []string) bool {
	keys := make(map[string]bool)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
		} else {
			return true
		}
	}
	return false
}
