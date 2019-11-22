package main

import (
	"container/heap"
	"fmt"
)

type HeapNode struct {
	Prime int
	Idx   int
	Val   int
}

type Heap struct {
	buf  []HeapNode
	size int
}

func NewHeap(n int) Heap {
	return Heap{make([]HeapNode, n), 0}
}

func (h Heap) Less(i, j int) bool {
	return h.buf[i].Val < h.buf[j].Val
}

func (h Heap) Swap(i, j int) {
	h.buf[i], h.buf[j] = h.buf[j], h.buf[i]
}

func (h Heap) Len() int {
	return h.size
}

func (h *Heap) Push(x interface{}) {
	h.buf[h.size] = x.(HeapNode)
	h.size++
}

func (h *Heap) Pop() interface{} {
	h.size--
	return h.buf[h.size]
}

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	h := NewHeap(len(primes))
	for _, p := range primes {
		heap.Push(&h, HeapNode{p, 0, p * dp[0]})
	}
	fmt.Println(h.buf)

	for i := 1; i < n; {
		//node := heap.Pop(&h).(HeapNode)
		node := &h.buf[0]
		if node.Val != dp[i-1] {
			dp[i] = node.Val
			i++
		}
		node.Idx++
		node.Val = node.Prime * dp[node.Idx]
		//heap.Push(&h, node)
		heap.Fix(&h, 0)
	}
	fmt.Println(dp)

	return dp[n-1]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}
