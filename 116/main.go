package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (n *Node) ToString() string {
	return ""
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	current := []*Node{root}
	next := []*Node{&Node{}} // &Node{} dummy head
	for len(current) > 0 {
		for _, node := range current {
			if node.Left != nil {
				next[len(next)-1].Next = node.Left
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next[len(next)-1].Next = node.Right
				next = append(next, node.Right)
			}
		}
		current = next[1:]
		next = []*Node{&Node{}}
	}
	return root
}

func main() {
	fmt.Println("vim-go")
}
