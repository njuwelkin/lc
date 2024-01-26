package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] == 0 {
		nums[0] = -1
	}
	m := map[int]int{0: -1, nums[0]: 0}
	ret := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		nums[i] += nums[i-1]
		if idx, ok := m[nums[i]]; ok {
			ret = max(ret, i-idx)
		} else {
			m[nums[i]] = i
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	fmt.Println(findMaxLength([]int{0, 0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0, 0, 0, 1, 0, 1})) // -1, 0, -1, -2, -3, -2, -3, -2
}
