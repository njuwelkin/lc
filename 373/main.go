package main

import (
	"container/heap"
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type sumNode struct {
	i, j int
	val  int
}

type sumHeap []sumNode

func (h sumHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h sumHeap) Len() int {
	return len(h)
}

func (h sumHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h *sumHeap) Push(node interface{}) {
	*h = append(*h, node.(sumNode))
}

func (h *sumHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1 := min(k, len(nums1))
	n2 := min(k, len(nums2))
	k = min(k, n1*n2)

	if n1 == 0 || n2 == 0 {
		return [][]int{}
	}

	ret := [][]int{}
	//h := sumHeap(make([]sumNode, 0, k))
	h := sumHeap([]sumNode{})

	for i := 0; i < n1; i++ {
		h.Push(sumNode{i, 0, nums1[i] + nums2[0]})
	}

	for k > 0 && h.Len() > 0 {
		node := heap.Pop(&h).(sumNode)
		ret = append(ret, []int{nums1[node.i], nums2[node.j]})

		i, j := node.i, node.j+1
		if j < n2 {
			node2 := sumNode{i, j, nums1[i] + nums2[j]}
			heap.Push(&h, node2)
		}
		k--
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}
