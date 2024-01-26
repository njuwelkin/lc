package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

type TreeIter struct {
	root  *TreeNode
	stack []*TreeNode
}

func NewTreeIter(root *TreeNode) TreeIter {
	stack := []*TreeNode{}
	for p := root; p != nil; p = p.Left {
		stack = append(stack, p)
	}
	return TreeIter{
		root:  root,
		stack: stack,
	}
}

func (ti TreeIter) Eof() bool {
	return len(ti.stack) == 0
}

func (ti *TreeIter) Get() int {
	p := ti.stack[len(ti.stack)-1]
	ti.stack = ti.stack[:len(ti.stack)-1]
	ret := p.Val
	for p = p.Right; p != nil; p = p.Left {
		ti.stack = append(ti.stack, p)
	}
	fmt.Println(ret)
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinimumDifference(root *TreeNode) int {
	ret := 10000000
	ti := NewTreeIter(root)
	prev := ti.Get()
	for !ti.Eof() {
		v := ti.Get()
		ret = min(ret, v-prev)
		prev = v
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(getMinimumDifference(NewBTree([]int{4, 2, 6, 1, 3})))
}
