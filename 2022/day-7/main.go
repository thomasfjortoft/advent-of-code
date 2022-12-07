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
	dirContent := getFolderContent(input)
	var total int

	for key := range dirContent {
		if out := getFolderSize(dirContent, key); out <= 100000 {
			total += out
		}
	}

	return total
}

func part2(input string) int {
	dirContent := getFolderContent(input)
	rootSize := getFolderSize(dirContent, "/")
	smallest := 100000000

	for key := range dirContent {
		if size := getFolderSize(dirContent, key); size >= 30000000-(70000000-rootSize) && size < smallest {
			smallest = size
		}
	}

	return smallest
}

func getFolderContent(input string) map[string][]string {
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
					tempDirList = nil
				}

				pathList = append(pathList, path)
				pathDir = strings.Join(pathList, "-")
			}
		} else {
			tempDirList = append(tempDirList, commandOrOutput)
		}
	}

	dirContent[pathDir] = tempDirList

	return dirContent
}

func getFolderSize(main map[string][]string, path string) int {
	var size int
	for _, item := range main[path] {
		splittedItem := strings.Split(item, " ")
		if strings.HasPrefix(splittedItem[0], "dir") {
			size += getFolderSize(main, path+"-"+splittedItem[1])
		} else {
			i, _ := strconv.Atoi(splittedItem[0])
			size += i
		}
	}
	return size
}
