package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func flatten(root *TreeNode) {
	prev := &TreeNode{}

	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		prev.Left = root
		prev = root
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	for p := root; p != nil; {
		p.Right = p.Left
		tmp := p.Left
		p.Left = nil
		p = tmp
	}
	fmt.Println(root.ToString())
	return
}

func main() {
	fmt.Println("vim-go")
	flatten(NewBTree([]int{1, 2, 5, 3, 4, Null, 6}))
}
