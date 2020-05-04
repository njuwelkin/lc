package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	var i, j int
	for i, j = 0, n; i < j; {

		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	start := search(nums, target)
	if start == n || nums[start] != target {
		return []int{-1, -1}
	}
	end := search(nums[start+1:], target+1) + start
	return []int{start, end}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 11))
}
