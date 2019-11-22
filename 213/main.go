package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// fn = max(f(n-1), f(n-2)+nums[n])
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	fn_2 := nums[0]
	fn_1 := max(nums[0], nums[1])
	fn := fn_1
	for i := 2; i < len(nums); i++ {
		fn = max(fn_1, fn_2+nums[i])
		fn_2 = fn_1
		fn_1 = fn
	}
	return fn
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}
func main() {
	fmt.Println("vim-go")
}
