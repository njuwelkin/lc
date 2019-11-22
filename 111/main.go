package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func minDepth(root *TreeNode) int {
	var dfs func(*TreeNode, int) bool
	dfs = func(root *TreeNode, limit int) bool {
		if limit == 0 {
			return false
		}
		if root.Left == nil && root.Right == nil {
			return true
		}
		if root.Left != nil && dfs(root.Left, limit-1) {
			return true
		}
		return root.Right != nil && dfs(root.Right, limit-1)
	}
	if root == nil {
		return 0
	}
	for i := 1; ; i++ {
		if dfs(root, i) {
			return i
		}
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minDepth(NewBTree([]int{3, 9, 20, Null, Null, 15, 7})))
}
