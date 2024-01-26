package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	prev := func(i int) int {
		if i == 0 {
			return math.MinInt
		}
		return nums[i-1]
	}
	next := func(i int) int {
		if i == len(nums)-1 {
			return math.MinInt
		}
		return nums[i+1]
	}

	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) / 2
		fmt.Println()
		if nums[mid] > prev(mid) && nums[mid] > next(mid) {
			return mid
		} else if nums[mid] <= prev(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1, 2, 3}))
}
