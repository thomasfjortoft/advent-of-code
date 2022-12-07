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

func part1(input string) int64 {
	dirContent := make(map[string][]string)
	var pathDir string
	var tempDirList []string

	var pathList []string

	for _, commandOrOutput := range strings.Split(input, "\n") {
		if strings.HasPrefix(commandOrOutput, "$") {

			var command, path string
			fmt.Sscanf(commandOrOutput, "$ %s %s", &command, &path)

			if path == ".." {

				pathList = pathList[:len(pathList)-1]
			} else if command == "cd" {
				if len(tempDirList) > 0 {
					dirContent[pathDir] = tempDirList
				}
				pathList = append(pathList, path)
				
				pathDir = strings.Join(pathList, "-")
				
			} else if command == "ls" {
				tempDirList = nil
			}

		} else {
			tempDirList = append(tempDirList, commandOrOutput)
		}
	}

	dirContent[pathDir] = tempDirList

	var total int64
	for key := range dirContent {
		if out := getFolderContent(dirContent, key); out <= 100000 {
			total += out
		}
	}

	return total
}

func part2(input string) int {
	return 0
}

func getFolderContent(main map[string][]string, path string) int64 {
	var size int64
	for _, item := range main[path] {
		if s := strings.Split(item, " "); strings.HasPrefix(s[0], "dir") {
			size += getFolderContent(main, path + "-" + strings.Split(item, " ")[1])
		} else {
			i, _ := strconv.ParseInt(strings.Split(item, " ")[0], 10, 64)
			size += i
		}
	}
	return size
}
