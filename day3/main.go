package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const fabricSize = 1000

var lineParser = regexp.MustCompile(`^#(.+)\s@\s(.+),(.+):\s(.+)x(.+)$`)

func main() {
	data, _ := ioutil.ReadFile("day3/input.txt")
	lines := strings.Split(string(data), "\n")

	fabric := make([][]int, fabricSize)

	for _, v := range lines {
		match := lineParser.FindStringSubmatch(strings.TrimSpace(v))
		_, _ = strconv.Atoi(match[1])
		left, _ := strconv.Atoi(match[2])
		top, _ := strconv.Atoi(match[3])
		width, _ := strconv.Atoi(match[4])
		height, _ := strconv.Atoi(match[5])

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if fabric[x+left] == nil {
					fabric[x+left] = make([]int, fabricSize)
				}
				fabric[x+left][y+top]++
			}
		}
	}

	moreThan := 0
	for _, row := range fabric {
		for _, col := range row {
			if col > 1 {
				moreThan++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", moreThan)

outer:
	for _, v := range lines {
		match := lineParser.FindStringSubmatch(strings.TrimSpace(v))
		id, _ := strconv.Atoi(match[1])
		left, _ := strconv.Atoi(match[2])
		top, _ := strconv.Atoi(match[3])
		width, _ := strconv.Atoi(match[4])
		height, _ := strconv.Atoi(match[5])

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if fabric[x+left][y+top] > 1 {
					continue outer
				}
			}
		}

		fmt.Printf("Part 2: ID #%d\n", id)
		break
	}

}
