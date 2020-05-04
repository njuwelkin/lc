package main

import "fmt"

func jump(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += i
		if nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
	}
	ret := 0
	for i := 0; i < len(nums)-1; i = nums[i] {
		ret++
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
