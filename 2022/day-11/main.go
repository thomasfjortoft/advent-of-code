package main

import (
	"fmt"
	"os"
	"sort"
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
	return monkeyWorryLevel(input, true, 20)
}

func part2(input string) int {
	return monkeyWorryLevel(input, false, 10000)
}

func monkeyWorryLevel(input string, worryLevelDivided bool, rounds int) int {

	listOfMonkeys := getMonkeyInput(input)
	var throws []int

	d := 1
	for _, mm := range listOfMonkeys {
		d *= mm.divisibleBy
	}

	for r := 0; r < rounds; r++ {

		for i := 0; i < len(listOfMonkeys); i++ {
			if r == 0 {
				throws = append(throws, 0)
			}
			currentMonkey := &listOfMonkeys[i]

			if len(currentMonkey.items) == 0 {
				continue
			}

			for _, item := range currentMonkey.items {

				throws[currentMonkey.name] = throws[currentMonkey.name] + 1

				currentMonkey.items = currentMonkey.items[1:]

				var worryLevel int64
				worryLevel = item
				if currentMonkey.arithmeticOperation == "*" {
					if currentMonkey.numberForOperation == 0 {
						worryLevel *= worryLevel
					} else {
						worryLevel *= currentMonkey.numberForOperation
					}
				} else if currentMonkey.arithmeticOperation == "+" {
					worryLevel += currentMonkey.numberForOperation
				}

				if worryLevelDivided {
					worryLevel = worryLevel / 3
				}

				worryLevel = worryLevel%int64(d)
				if worryLevel%int64(currentMonkey.divisibleBy) == 0 {
					listOfMonkeys[currentMonkey.throwToIfTrue].items = append(listOfMonkeys[currentMonkey.throwToIfTrue].items, worryLevel)
				} else {
					listOfMonkeys[currentMonkey.throwToIfFalse].items = append(listOfMonkeys[currentMonkey.throwToIfFalse].items, worryLevel)
				}
			}

		}
	}

	sort.Ints(throws)

	return throws[len(throws)-1] * throws[len(throws)-2]
}

type monkey struct {
	name                int
	items               []int64
	arithmeticOperation string
	numberForOperation  int64
	divisibleBy         int
	throwToIfTrue       int
	throwToIfFalse      int
}

func getMonkeyInput(input string) []monkey {
	listOfMonkeys := []monkey{}

	for _, monkeyText := range strings.Split(input, "\n\n") {
		monkeyInput := strings.Split(monkeyText, "\n")
		var (
			name                int
			arithmeticOperation string
			numberForOperation  int64
			divisibleBy         int
			throwToIfTrue       int
			throwToIfFalse      int
		)
		fmt.Sscanf(monkeyInput[0], "Monkey %d:", &name)
		strings := strings.Split(strings.Split(monkeyInput[1], ": ")[1], ", ")
		ints := make([]int64, len(strings))

		for i, s := range strings {
			t, _ := strconv.Atoi(s)
			ints[i] = int64(t)
		}
		fmt.Sscanf(monkeyInput[2], "  Operation: new = old %s %d", &arithmeticOperation, &numberForOperation)
		fmt.Sscanf(monkeyInput[3], "  Test: divisible by %d", &divisibleBy)
		fmt.Sscanf(monkeyInput[4], "    If true: throw to monkey %d", &throwToIfTrue)
		fmt.Sscanf(monkeyInput[5], "    If false: throw to monkey %d", &throwToIfFalse)

		listOfMonkeys = append(listOfMonkeys, monkey{
			name:                name,
			items:               ints,
			arithmeticOperation: arithmeticOperation,
			numberForOperation:  numberForOperation,
			divisibleBy:         divisibleBy,
			throwToIfTrue:       throwToIfTrue,
			throwToIfFalse:      throwToIfFalse,
		})
	}
	return listOfMonkeys
}
