package main

import "fmt"

func findMin(nums []int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] < nums[j-1] {
			return nums[i]
		}
		mid := (i + j) / 2
		if nums[mid] > nums[i] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return nums[i]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
