package main

import (
	"container/heap"
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

type node struct {
	sum   int
	count int
}

type Heap []node

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h Heap) Push(x interface{}) {
	h = append(h, x.(node))
}

func (h Heap) Pop() interface{} {
	ret := h[len(h)-1]
	h = h[:len(h)-1]
	return ret
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	count := map[int]int{}

	var postOrder func(*TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		crt := root.Val + postOrder(root.Left) + postOrder(root.Right)
		if c, ok := count[crt]; ok {
			count[crt] = c + crt
		} else {
			count[crt] = crt
		}
		return crt
	}

	postOrder(root)

	h := Heap{}
	heap.Init(h)

	for sum, c := range count {
		heap.Push(h, node{sum, c})
	}

	n := heap.Pop(h).(node)
	c := n.count

	ret := []int{n.sum}
	for h.Len() > 0 {
		n = heap.Pop(h).(node)
		if n.count == c {
			ret = append(ret, n.sum)
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}
