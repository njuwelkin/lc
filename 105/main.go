package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 || len(inorder) != n {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := make([]*TreeNode, n)
	stack[0] = root
	fmt.Println("push ", preorder[0])
	top := 1

	prevPoint := &root.Left

	var i, j int
	for i, j = 1, 0; j < n; {
		if top != 0 && stack[top-1].Val == inorder[j] {
			fmt.Println("pop ", inorder[j])
			top--
			j++
			prevPoint = &stack[top].Right
		} else {
			if i == n {
				panic("")
			}
			fmt.Println("push ", preorder[i])
			stack[top] = new(TreeNode)
			stack[top].Val = preorder[i]
			*prevPoint = stack[top]
			prevPoint = &stack[top].Left
			top++
			i++
		}
	}

	if j != n || i != n || top != 0 {
		fmt.Println("error input", i, j, top)
		return nil
	}
	fmt.Println(root)
	return root
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).ToString())
}
