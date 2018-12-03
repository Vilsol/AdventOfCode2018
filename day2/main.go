package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("day2/input.txt")
	lines := strings.Split(string(data), "\n")

	twos := 0
	threes := 0

	for _, line := range lines {
		chars := make(map[int32]int)
		for _, char := range line {
			chars[char] = chars[char] + 1
		}

		foundTwo := false
		foundThree := false

		for _, count := range chars {
			if count == 2 && !foundTwo {
				twos++
				foundTwo = true
			}
			if count == 3 && !foundThree {
				threes++
				foundThree = true
			}
			if foundTwo && foundThree {
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", twos*threes)

outer:
	for _, line := range lines {
	inner:
		for _, line2 := range lines {
			different := 0
			for i := range line2 {
				if line[i] != line2[i] {
					if different >= 1 {
						continue inner
					}
					different++
				}
			}

			if different == 1 {
				result := ""
				for i := range line2 {
					if line[i] == line2[i] {
						result = result + string(line2[i])
					}
				}
				fmt.Printf("Part 2: %s\n", result)
				break outer
			}
		}
	}

}
