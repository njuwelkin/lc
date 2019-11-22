package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPathSum(root *TreeNode) int {
	ret := math.MinInt64
	var postOrder func(root *TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(postOrder(root.Left), 0)
		r := max(postOrder(root.Right), 0)
		ret = max(ret, root.Val+l+r)
		return max(l, r) + root.Val
	}
	postOrder(root)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxPathSum(NewBTree([]int{-10, 9, 20, Null, Null, 15, 7})))
}
