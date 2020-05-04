package main

import "fmt"

func canJump(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		nums[i] += i
		if nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
		if nums[i] == i {
			return false
		}
		if nums[i] >= len(nums)-1 {
			return true
		}
	}
	fmt.Println(nums)
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
