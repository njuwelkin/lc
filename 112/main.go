package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func hasPathSum1(root *TreeNode, sum int) bool {
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return root.Left != nil && hasPathSum1(root.Left, sum-root.Val) ||
		root.Right != nil && hasPathSum1(root.Right, sum-root.Val)
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	}
	return hasPathSum1(root, sum)
}

func main() {
	fmt.Println("vim-go")
}
