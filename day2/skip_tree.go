package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Children map[int32]*Node
}

type NodeStr struct {
	Children map[string]*NodeStr
}

func (n NodeStr) Dump(indent int) {
	for k, v := range n.Children {
		fmt.Printf("%s %s\n", strings.Repeat("-", indent), k)
		v.Dump(indent + 1)
	}
}
