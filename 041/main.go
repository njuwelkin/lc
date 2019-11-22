package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	for _, num := range nums {
		for j := num - 1; j >= 0 && j < len(nums) && j != nums[j]-1; {
			tmp := nums[j]
			nums[j] = j + 1
			j = tmp - 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
