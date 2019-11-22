package main

import "fmt"

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	if k >= len(nums) {
		k %= len(nums)
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

func main() {
	fmt.Println("vim-go")
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
