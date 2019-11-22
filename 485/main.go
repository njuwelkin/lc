package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func findMaxConsecutiveOnes(nums []int) int {
	ret := 0
	n := len(nums)
	var i, j int
	for i = 0; i < n; {
		for ; i < n && nums[i] == 0; i++ {
		}

		for j = i; j < n && nums[j] == 1; j++ {
		}
		ret = max(ret, j-i)
		i = j
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
