package main

import (
	"fmt"
	"math"
)

type Seg struct {
	min, max int
}

func maximumGap(nums []int) int {
	min, max := math.MaxInt64, math.MinInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	n := len(nums)
	scope := max + 1 - min
	segLen := (scope + n) / (n + 1)
	segs := make([]Seg, n+1)
	for i := 0; i < n+1; i++ {
		segs[i].min = math.MaxInt64
		segs[i].max = math.MinInt64
	}
	for _, num := range nums {
		idx := (num - min) / segLen
		if num > segs[idx].max {
			segs[idx].max = num
		}
		if num < segs[idx].min {
			segs[idx].min = num
		}
	}

	ret := 0
	prevSeg := 0
	for i := 1; i < n+1; i++ {
		if segs[i].min == math.MaxInt64 {
			continue
		}
		if segs[i].min-segs[prevSeg].max > ret {
			ret = segs[i].min - segs[prevSeg].max
		}
		prevSeg = i
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
}
