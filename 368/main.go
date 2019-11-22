package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	sort.Ints(nums)
	depth := make([]int, n)
	prev := make([]int, n)
	var i, j int

	maxIdx := 0
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if depth[j] < depth[i]+1 && nums[j]%nums[i] == 0 {
				depth[j] = depth[i] + 1
				prev[j] = i + 1
				if depth[j] > depth[maxIdx] {
					maxIdx = j
				}
			}
		}
	}

	ret := []int{}
	for i = maxIdx; i >= 0; i = prev[i] - 1 {
		ret = append(ret, nums[i])
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(largestDivisibleSubset([]int{}))
	fmt.Println(largestDivisibleSubset([]int{1}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 9}))
}
