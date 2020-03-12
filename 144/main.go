package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	ret := []int{}
	p := root

	for {
		for p != nil {
			ret = append(ret, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) == 0 {
			return ret
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			p = top.Right
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(NewBTree([]int{1, Null, 2, 3}).ToString())
	fmt.Println(preorderTraversal(NewBTree([]int{1, Null, 2, 3})))
	fmt.Println(preorderTraversal(nil))
}
