package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

type BTIter struct {
	stack []*TreeNode
	root  *TreeNode
}

func NewBTIter(root *TreeNode) *BTIter {
	return (&BTIter{[]*TreeNode{}, root}).Reset()
}

func (bti *BTIter) Reset() *BTIter {
	bti.stack = []*TreeNode{}
	p := bti.root
	for p != nil {
		bti.stack = append(bti.stack, p)
		p = p.Left
	}
	return bti
}

func (bti *BTIter) Next() *TreeNode {
	if len(bti.stack) == 0 {
		return nil
	}
	top := bti.stack[len(bti.stack)-1]
	bti.stack = bti.stack[:len(bti.stack)-1]
	p := top.Right
	for p != nil {
		bti.stack = append(bti.stack, p)
		p = p.Left
	}
	return top
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	bti1, bti2 := NewBTIter(root1), NewBTIter(root2)
	p1, p2 := bti1.Next(), bti2.Next()
	ret := []int{}

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			ret = append(ret, p1.Val)
			p1 = bti1.Next()
		} else {
			ret = append(ret, p2.Val)
			p2 = bti2.Next()
		}
	}

	for p1 != nil {
		ret = append(ret, p1.Val)
		p1 = bti1.Next()
	}
	for p2 != nil {
		ret = append(ret, p2.Val)
		p2 = bti2.Next()
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(getAllElements(NewBTree([]int{2, 1, 4}), NewBTree([]int{1, 0, 3})))
	fmt.Println(getAllElements(NewBTree([]int{1, Null, 8}), NewBTree([]int{8, 1})))
}
