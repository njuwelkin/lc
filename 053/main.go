package main

import "fmt"

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}

func maxSubArray(nums []int) int {
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
