package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

func convertBST(root *TreeNode) *TreeNode {
	prev := 0
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Right)
		prev += root.Val
		root.Val = prev
		f(root.Left)
	}
	f(root)
	return root
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(convertBST(NewBTree([]int{4, 1, 6, 0, 2, 5, 7, Null, Null, Null, 3, Null, Null, Null, 8})).ToString())
}
