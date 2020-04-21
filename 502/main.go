package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func findMaxIdx(W int, Profits []int, Capital []int) int {
	maxIdx := -1
	maxVal := -1 << 63
	for i := 0; i < len(Capital); i++ {
		if Capital[i] <= W && Profits[i] > maxVal {
			maxIdx = i
			maxVal = Profits[i]
		}
	}
	return maxIdx
}

func findMaximizedCapital2(k int, W int, Profits []int, Capital []int) int {
	for i := 0; i < k; i++ {
		fmt.Println(Profits, Capital)
		idx := findMaxIdx(W, Profits, Capital)
		fmt.Println("idx:", idx)
		if idx < 0 {
			break
		}
		W += Profits[idx]
		fmt.Println("W:", W)
		Profits[idx] = Profits[len(Profits)-1]
		Profits = Profits[:len(Profits)-1]
		Capital[idx] = Capital[len(Capital)-1]
		Capital = Capital[:len(Capital)-1]
	}
	return W
}

type ProfitHeap []int

func (this *ProfitHeap) Len() int {
	return len(*this)
}

func (this *ProfitHeap) Less(i, j int) bool {
	return (*this)[i] > (*this)[j]
}

func (this *ProfitHeap) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *ProfitHeap) Push(x interface{}) {
	val, _ := x.(int)
	(*this) = append(*this, val)
}

func (this *ProfitHeap) Pop() interface{} {
	ret := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return ret
}

func find(buf [][]int, capital int) int {
	i, j := 0, len(buf)
	for i < j {
		mid := (i + j) / 2
		if buf[mid][0] <= capital {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	buf := make([][]int, len(Profits))
	for i := range Profits {
		buf[i] = []int{Capital[i], Profits[i]}
	}
	sort.Slice(buf, func(i, j int) bool {
		return buf[i][0] < buf[j][0] || buf[i][0] == buf[j][0] && buf[i][1] > buf[j][1]
	})

	pHeap := ProfitHeap(Profits[:0])

	i := 0
	for ; k > 0; k-- {
		/*
			end := i + find(buf[i:], W)
			for ; i < end; i++ {
				heap.Push(&pHeap, buf[i][1])
			}
		*/
		for ; i < len(buf) && buf[i][0] <= W; i++ {
			heap.Push(&pHeap, buf[i][1])
		}
		if pHeap.Len() == 0 {
			break
		}
		max, _ := heap.Pop(&pHeap).(int)
		W += max
	}

	return W
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))
}
