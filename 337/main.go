package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Val < 0 {
		return -root.Val
	}
	a := calc(root.Left) + calc(root.Right)
	b := root.Val
	if root.Left != nil {
		b += calc(root.Left.Left) + calc(root.Left.Right)
	}
	if root.Right != nil {
		b += calc(root.Right.Left) + calc(root.Right.Right)
	}
	root.Val = -max(a, b)
	return -root.Val
}

func main() {
	fmt.Println("vim-go")
}
