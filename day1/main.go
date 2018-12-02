package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("day1/input.txt")
	lines := strings.Split(string(data), "\n")

	result := 0
	twice := false
	first := true

	visited := make(map[int]bool)
	visited[0] = true

	for !twice {
		for _, line := range lines {
			i := 0

			if line[0] == '+' {
				i, _ = strconv.Atoi(strings.TrimSpace(line[1:]))
			} else {
				i, _ = strconv.Atoi(strings.TrimSpace(line[1:]))
				i *= -1
			}

			result += i

			if visited[result] {
				twice = true
				fmt.Printf("Visited twice: %d\n", result)

				if !first {
					break
				}
			}

			visited[result] = true
		}

		if first {
			first = false
			fmt.Printf("Resulting frequency: %d\n", result)
		}
	}
}