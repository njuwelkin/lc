package main

import "fmt"

func minPatches(nums []int, n int) int {
	count := 0
	j := 0
	for i := 1; i <= n; {
		if j >= len(nums) || i < nums[j] {
			count++
			fmt.Print(i, " ")
			i *= 2
		} else {
			i += nums[j]
			j++
		}
	}
	return count
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(minPatches([]int{1, 3}, 6))
	//fmt.Println(minPatches([]int{1, 5, 10}, 20))
	//fmt.Println(minPatches([]int{1, 2, 2}, 5))
	fmt.Println(minPatches([]int{}, 8))
}
