package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func serialize(root *TreeNode) []int {
	ret := []int{}

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ret
}

func deserialize(data []int) *TreeNode {
	dummyRoot := &TreeNode{Val: 1 << 62}
	stack := []*TreeNode{dummyRoot}

	var prev *TreeNode
	for i := 0; i < len(data); {
		top := stack[len(stack)-1]
		if data[i] < top.Val {
			node := &TreeNode{data[i], nil, nil}
			if top.Left == nil {
				top.Left = node
			} else {
				prev.Right = node
			}
			stack = append(stack, node)
			i++
		} else {
			stack = stack[:len(stack)-1]
			prev = top
		}
	}
	return dummyRoot.Left
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(deserialize([]int{8, 3, 1, 6, 4, 7, 10, 14, 13}).ToString())
	fmt.Println(serialize(deserialize([]int{8, 3, 1, 6, 4, 7, 10, 14, 13})))
}
