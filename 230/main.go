package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func kthSmallest(root *TreeNode, k int) int {
	ret := 0
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if dfs(root.Left) {
			return true
		}
		k--
		if k == 0 {
			ret = root.Val
			return true
		}
		return dfs(root.Right)
	}
	dfs(root)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(kthSmallest(NewBTree([]int{3, 1, 4, Null, 2}), 1))
	fmt.Println(kthSmallest(NewBTree([]int{5, 3, 6, 2, 4, Null, Null, 1}), 3))
}
