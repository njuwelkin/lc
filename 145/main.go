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

type stackNode struct {
	p    *TreeNode
	stat int
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*stackNode{&stackNode{root, 0}}
	ret := []int{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top.stat == 0 {
			if top.p.Left != nil {
				stack = append(stack, &stackNode{top.p.Left, 0})
			}
		} else if top.stat == 1 {
			if top.p.Right != nil {
				stack = append(stack, &stackNode{top.p.Right, 0})
			}
		} else {
			ret = append(ret, top.p.Val)
			stack = stack[:len(stack)-1]
		}
		top.stat++
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(NewBTree([]int{1, Null, 2, 3}).ToString())
	fmt.Println(postorderTraversal(NewBTree([]int{1, Null, 2, 3})))
	fmt.Println(postorderTraversal(nil))
}
