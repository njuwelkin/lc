package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

func findBottomLeftValue(root *TreeNode) int {
	maxDepth := 0
	ret := 0

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > maxDepth {
			ret = root.Val
			maxDepth = depth
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 1)

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findBottomLeftValue(NewBTree([]int{2, 1, 3})))
	fmt.Println(findBottomLeftValue(NewBTree([]int{1, 2, 3, 4, Null, 5, 6, Null, Null, 7})))
}
