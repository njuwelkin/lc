package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var i int
	for i = 1; i < len(nums) && nums[i] == nums[0]; i++ {
	}
	if i == len(nums) {
		return 1
	}
	prev := nums[i] - nums[0]

	ret := 2
	for i = i + 1; i < len(nums); i++ {
		if (nums[i]-nums[i-1])*prev >= 0 {
			continue
		}
		prev = nums[i] - nums[i-1]
		ret++
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}
