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

func diameterOfBinaryTree(root *TreeNode) int {
	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		diameter := 0
		depthL, depthR := 0, 0
		if root.Left != nil {
			diameter = f(root.Left)
			depthL = root.Left.Val + 1
		}
		if root.Right != nil {
			diameter = max(diameter, f(root.Right))
			depthR = root.Right.Val + 1
		}
		root.Val = max(depthL, depthR)
		return max(diameter, depthL+depthR)
	}
	if root != nil {
		ret := f(root)
		fmt.Println(root.ToString())
		return ret
	}
	return 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(diameterOfBinaryTree(NewBTree([]int{1, 2, 3, 4, 5})))
	fmt.Println(diameterOfBinaryTree(NewBTree([]int{1, 2})))
}
