package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	crt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[crt] = nums[i]
			crt++
		}
	}
	return crt
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
