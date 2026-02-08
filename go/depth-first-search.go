package main

import "fmt"

type Node struct {
	value    int
	children []*Node
}

func dfs(curr *Node, needle int) bool {
	if curr == nil {
		return false
	}
	if curr.value == needle {
		return true
	}
	for _, child := range curr.children {
		if dfs(child, needle) {
			return true
		}
	}
	return false
}

func main() {
	root := &Node{
		value: 10,
		children: []*Node{
			{value: 7},
			{
				value: 15,
				children: []*Node{
					{value: 12},
				},
			},
			{value: 20},
		},
	}
	fmt.Println(dfs(root, 20))
	fmt.Println(dfs(root, 21))
}
