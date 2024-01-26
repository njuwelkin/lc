package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	i, j := 0, (len(nums)+1)/2
	for i < j {
		mid := (i + j) / 2
		if 2*mid+1 < len(nums) && nums[2*mid] == nums[2*mid+1] {
			i = mid + 1
		} else {
			if mid == 0 || nums[2*mid] != nums[2*mid-1] {
				return nums[2*mid]
			}
			j = mid
		}
	}
	return nums[i*2]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 3, 3, 4, 4, 8, 8, 9}))
	fmt.Println(singleNonDuplicate([]int{1}))
}
