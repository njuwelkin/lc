package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Heap struct {
	sort.IntSlice
	IsReverse bool
}

func NewHeap(reverse bool) Heap {
	return Heap{[]int{}, reverse}
}

func (h Heap) Less(i, j int) bool {
	if h.IsReverse {
		return h.IntSlice.Less(j, i)
	}
	return h.IntSlice.Less(i, j)
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	ret := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return ret
}

type MedianFinder struct {
	maxHeap Heap
	minHeap Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	ret := MedianFinder{}
	ret.minHeap = NewHeap(false)
	ret.maxHeap = NewHeap(true)
	return ret
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() == 0 {
		heap.Push(&this.minHeap, num)
		return
	}
	if num > this.minHeap.IntSlice[0] {
		if this.minHeap.Len() > this.maxHeap.Len() {
			tmp := this.minHeap.IntSlice[0]
			this.minHeap.IntSlice[0] = num
			heap.Fix(&this.minHeap, 0)
			heap.Push(&this.maxHeap, tmp)
		} else {
			heap.Push(&this.minHeap, num)
		}
	} else if this.maxHeap.Len() == 0 || num < this.maxHeap.IntSlice[0] {
		if this.maxHeap.Len() == this.minHeap.Len() {
			tmp := this.maxHeap.IntSlice[0]
			this.maxHeap.IntSlice[0] = num
			heap.Fix(&this.maxHeap, 0)
			heap.Push(&this.minHeap, tmp)
		} else {
			heap.Push(&this.maxHeap, num)
		}
	} else {
		if this.maxHeap.Len() == this.minHeap.Len() {
			heap.Push(&this.minHeap, num)
		} else {
			heap.Push(&this.maxHeap, num)
		}

	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == 0 {
		return 0.0
	}
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.maxHeap.IntSlice[0]+this.minHeap.IntSlice[0]) / 2
	}
	return float64(this.minHeap.IntSlice[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor()
	//obj.AddNum(40)
	obj.AddNum(1)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	//obj.AddNum(12)
	obj.AddNum(2)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	//obj.AddNum(16)
	obj.AddNum(3)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	return
	obj.AddNum(14)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(35)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(19)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(34)
	obj.AddNum(35)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
	obj.AddNum(28)
	obj.AddNum(35)
	fmt.Println(obj.maxHeap, obj.minHeap)
	fmt.Println(obj.FindMedian())
}
