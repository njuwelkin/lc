package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	b := max(p.Val, q.Val)
	a := p.Val + q.Val - b

	for {
		if root == nil {
			panic("")
		}
		if root.Val >= a && root.Val <= b {
			return root
		}
		if root.Val < a {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

func main() {
	fmt.Println("vim-go")
	root := NewBTree([]int{6, 2, 8, 0, 4, 7, 9, Null, Null, 3, 5})
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).ToString())
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right).ToString())
}
