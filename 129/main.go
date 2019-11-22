package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func sumNumbers(root *TreeNode) int {
	ret := 0

	var f func(*TreeNode, int)
	f = func(root *TreeNode, val int) {
		if root.Left == nil && root.Right == nil {
			ret += val*10 + root.Val
			return
		}
		val = val*10 + root.Val
		if root.Left != nil {
			f(root.Left, val)
		}
		if root.Right != nil {
			f(root.Right, val)
		}
	}

	if root == nil {
		return 0
	}
	f(root, 0)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(sumNumbers(NewBTree([]int{1, 2, 3})))
	fmt.Println(sumNumbers(NewBTree([]int{4, 9, 0, 5, 1})))
}
