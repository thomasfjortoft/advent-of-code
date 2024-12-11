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
	startStones := parseInput(input)
	return blinking(startStones, 25)
}

func part2(input string) int {
	startStones := parseInput(input)
	return blinking(startStones, 75)
}

func blinking(startList []int, numberOfBlinks int) int {

	var count int
	for _, stone := range startList {
		count += blinkCount(stone, numberOfBlinks)
	}
	return count
}
var mem = make(map[[2]int]int)

func blinkCount(stone int, steps int) int {
	if v, ok := mem[[2]int{stone, steps}]; ok {
		return v
	}

	next := checkOneStone(stone)
	mem[[2]int{stone, 1}] = len(next)
	if steps == 1 {
		return len(next)
	}

	var count int
	for _, v := range next {
		count += blinkCount(v, steps-1)
	}
	mem[[2]int{stone, steps}] = count
	return count
}

func checkOneStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	} 
	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		return []int{toInt(s[:len(s)/2]), toInt(s[len(s)/2:])}
	} 
	return []int{stone * 2024}
}

func parseInput(input string) []int {
	parts := strings.Split(strings.TrimSpace(input), " ")
	nums := make([]int, len(parts))
	for i, part := range parts {
		nums[i] = toInt(part)
	}
	return nums
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
