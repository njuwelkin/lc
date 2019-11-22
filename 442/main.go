package main

import "fmt"

func findDuplicates(nums []int) []int {
	var i, j int
	mask := 1 << 61
	for i = 0; i < len(nums); i++ {
		val := ((mask - 1) & nums[i]) - 1
		fmt.Printf("%b", val)
		nums[val] += mask
	}
	fmt.Println(nums)

	mask <<= 1
	for i, j = 0, 0; j < len(nums); j++ {
		if nums[j]&mask != 0 {
			nums[i] = j + 1
			i++
		}
	}

	return nums[:i]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
