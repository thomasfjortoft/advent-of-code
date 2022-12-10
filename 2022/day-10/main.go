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
	fmt.Printf("Part2: %s took %s \n", part2(string(input)), time.Since(start))
}

func part1(input string) int {
	total, cycle, x := 0, 1, 1

	for _, v := range strings.Split(input, "\n") {
		if command := strings.Split(v, " "); command[0] == "addx" {
			//fmt.Println("addx")
			for i := 1; i <= 2; i++ {
				//fmt.Println("addx ", i)
				d, _ := strconv.Atoi(command[1])
				total += getSignalStrengthPoints(cycle, x)
				if i == 2 {
					x += d
				}
				cycle += 1
			}
		} else {
			total += getSignalStrengthPoints(cycle, x)
			cycle += 1
		}

	}

	return total
}

func getSignalStrengthPoints(cycle, x int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * x
	}
	return 0
}

func part2(input string) string {
	cycle, x, sprite := 1, 1, 0
	var crt []string
	row := ""
	for _, v := range strings.Split(input, "\n") {
		if command := strings.Split(v, " "); command[0] == "addx" {
			//fmt.Println("addx")
			for i := 1; i <= 2; i++ {
				//fmt.Println("addx ", i)
				if (cycle-1)%40 == 0 {
					crt = append(crt, row)
					row = ""
					sprite = 0
				}
				d, _ := strconv.Atoi(command[1])
				if sprite == x || sprite-1 == x || sprite+1 == x {
					row += "#"
				} else {
					row += "."
				}
				if i == 2 {
					x += d
				}
				cycle += 1
				sprite += 1
			}
		} else {
			if (cycle-1)%40 == 0 {
				crt = append(crt, row)
				row = ""
				sprite = 0
			}
			if sprite == x || sprite-1 == x || sprite+1 == x {
				row += "#"
			} else {
				row += "."
			}
			cycle += 1
			sprite += 1
		}

	}


	return strings.Join(crt, "") + row
}
