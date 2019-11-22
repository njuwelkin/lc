package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(invertTree(NewBTree([]int{4, 2, 7, 1, 3, 6, 9})).ToString())
}
