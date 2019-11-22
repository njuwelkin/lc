package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func search(prevPointer **TreeNode, node *TreeNode, key int) (**TreeNode, *TreeNode) {
	if node == nil {
		return prevPointer, nil
	}
	if node.Val == key {
		return prevPointer, node
	} else if key < node.Val {
		return search(&node.Left, node.Left, key)
	} else {
		return search(&node.Right, node.Right, key)
	}
}

func del(prevPointer **TreeNode, node *TreeNode) {
	if node.Left == nil && node.Right == nil {
		*prevPointer = nil
	} else if node.Left == nil {
		*prevPointer = node.Right
	} else if node.Right == nil {
		*prevPointer = node.Left
	} else {
		prevP := &node.Left
		p := node.Left
		for p.Right != nil {
			prevP = &p.Right
			p = p.Right
		}
		node.Val = p.Val
		del(prevP, p)
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	fakeRoot := TreeNode{}
	fakeRoot.Left = root
	prevP, node := search(&fakeRoot.Left, root, key)
	if node == nil {
		return root
	}
	del(prevP, node)
	return fakeRoot.Left
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(deleteNode(NewBTree([]int{5, 3, 6, 2, 4, Null, 7}), 3).Serialize())
}
