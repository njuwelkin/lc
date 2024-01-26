package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	current := []*TreeNode{root}
	for len(current) != 0 {
		next := []*TreeNode{}
		max := current[0].Val
		for _, p := range current {
			if p.Val > max {
				max = p.Val
			}
			if p.Left != nil {
				next = append(next, p.Left)
			}
			if p.Right != nil {
				next = append(next, p.Right)
			}
		}
		ret = append(ret, max)
		current = next
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(largestValues(NewBTree([]int{1, 3, 2, 5, 3, Null, 9})))
	fmt.Println(largestValues(NewBTree([]int{1, 2, 3})))
}
