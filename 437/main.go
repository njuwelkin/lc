package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func pathSumFromRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	match := 0
	if sum == root.Val {
		match = 1
	}
	return match + pathSumFromRoot(root.Left, sum-root.Val) + pathSumFromRoot(root.Right, sum-root.Val)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, sum) + pathSum(root.Right, sum) + pathSumFromRoot(root, sum)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(pathSum(NewBTree([]int{10, 5, -3, 3, 2, Null, 11, 3, -2, Null, 1}), 8))
	fmt.Println(pathSum(NewBTree([]int{}), 0))
}
