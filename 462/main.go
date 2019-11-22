package main

import (
	"fmt"
)

func trail(nums []int) int {
	n := len(nums)
	tmp := nums[0]
	var i, j int
	for i, j = 1, n-1; i <= j; {
		if nums[i] <= tmp {
			nums[i-1] = nums[i]
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	nums[i-1] = tmp
	return i - 1
}

func getKthNumber(nums []int, k int) int {
	if len(nums) == 1 && k == 0 {
		return nums[0]
	}
	i := trail(nums)
	if i < k {
		return getKthNumber(nums[i+1:], k-i-1)
	} else {
		return getKthNumber(nums[:i+1], k)
	}
}

func getMedian(nums []int) int {
	n := len(nums)
	return getKthNumber(nums, n/2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMoves2(nums []int) int {
	mid := getMedian(nums)
	ret := 0
	for _, num := range nums {
		ret += abs(num - mid)
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minMoves2([]int{1, 2, 3}))
	fmt.Println(minMoves2([]int{1, 0, 0, 8, 6}))
}
