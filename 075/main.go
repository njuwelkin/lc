package main

import "fmt"

func sortColors(nums []int) {
	var i, j, k int
	for i, j, k = 0, 0, len(nums)-1; j <= k; {
		if nums[j] == 0 {
			nums[i] = 0
			i++
			j++
		} else if nums[j] == 2 {
			nums[j] = nums[k]
			nums[k] = 2
			k--
		} else {
			j++
		}
	}

	for ; i <= k; i++ {
		nums[i] = 1
	}
	fmt.Println(nums)
}

func main() {
	fmt.Println("vim-go")
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
