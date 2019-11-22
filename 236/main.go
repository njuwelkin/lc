package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func search(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p || root == q {
		return true
	}
	return search(root.Left, p, q) || search(root.Right, p, q)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ret *TreeNode
	var search2 func(root *TreeNode) int
	search2 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root == p || root == q {
			if search(root.Left, p, q) || search(root.Right, p, q) {
				ret = root
				return 2
			}
			return 1
		} else {
			var left int
			if left = search2(root.Left); left == 2 {
				return 2
			} else if left == 1 {
				if search(root.Right, p, q) {
					ret = root
					return 2
				} else {
					return 1
				}
			} else {
				return search2(root.Right)
			}
		}
	}
	search2(root)
	return ret
}

func main() {
	fmt.Println("vim-go")
	root := NewBTree([]int{3, 5, 1, 6, 2, 0, 8, Null, Null, 7, 4})
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).Val)
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right).Val)
}
