package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const correctAnswer = "cnjxoritzhvbosyewrmqhgkul"

// const correctAnswer = "fgij"

var lines []string

func TestMain(m *testing.M) {
	data, _ := ioutil.ReadFile("input.txt")
	lines = strings.Split(string(data), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	os.Exit(m.Run())
}

func Unoptimized() string {
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
				return result
			}
		}
	}
	return ""
}

func Optimized() string {
	dataTree := &Node{make(map[int32]*Node)}
	for _, line := range lines {
		lastNode := dataTree

		searchNodes := make(map[int32]*Node)
		broken := 0

		for i, char := range line {

			/*
				if line == "cnjxoritdzhvbosyewrmqhgkul" && char == 'd' {
					fmt.Println()
					fmt.Println(line)
					fmt.Println(string(char))
					fmt.Println(line[i:])
					fmt.Println(len(searchNodes))
					for lk := range searchNodes {
						fmt.Println(lk)
					}
					fmt.Println(lastNode.Children)
				}
			*/

			if searchNodes != nil {
			search:
				for _, v := range searchNodes {
					searchNode := v
					for _, fchar := range line[i+1:] {
						/*
							if line == "cnjxoritdzhvbosyewrmqhgkul" && char == 'd' {
								fmt.Println("NEST:")
								fmt.Println(string(incorrect))
								for lk := range v.Children {
									fmt.Println(lk)
								}
							}
						*/
						if searchNode == nil {
							continue search
						}

						if searchNode.Children == nil ||
							searchNode.Children[fchar] == nil {
							continue search
						}
						searchNode = searchNode.Children[fchar]
					}

					return line[:i] + line[i+1:]
				}
				searchNodes = nil
			}

			if lastNode.Children[char] == nil {
				broken++
				lastNode.Children[char] = &Node{make(map[int32]*Node)}
			}

			lastNode = lastNode.Children[char]

			if broken <= 1 {
				searchNodes = lastNode.Children
			}
		}
	}
	return ""
}

func TestUnoptimized(t *testing.T) {
	result := Unoptimized()
	if result != correctAnswer {
		t.Errorf("Incorrect Answer. Expected %s. Got %s", correctAnswer, result)
	}
}

func BenchmarkUnoptimized(b *testing.B) {
	for i := 0; i < 1000; i++ {
		Unoptimized()
	}
}

func TestOptimized(t *testing.T) {
	result := Optimized()
	if result != correctAnswer {
		t.Errorf("Incorrect Answer. Expected %s. Got %s", correctAnswer, result)
	}
}

func BenchmarkOptimized(b *testing.B) {
	for i := 0; i < 1000; i++ {
		Optimized()
	}
}
