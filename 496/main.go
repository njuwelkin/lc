package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := map[int]int{}
	stack := make([]int, len(nums2))
	top := 0
	for _, num := range nums2 {
		for top > 0 && stack[top-1] < num {
			greater[stack[top-1]] = num
			top--
		}
		stack[top] = num
		top++
	}

	ret := make([]int, len(nums1))
	for i, num := range nums1 {
		var found bool
		ret[i], found = greater[num]
		if !found {
			ret[i] = -1
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
