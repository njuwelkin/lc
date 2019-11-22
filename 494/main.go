package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	prev := map[int]int{nums[0]: 2}
	var current map[int]int
	for i := 1; i < n; i++ {
		current = make(map[int]int)
		for sum, count := range prev {
			current[abs(sum+nums[i])] += count / 2
			current[abs(sum-nums[i])] += count / 2
			current[abs(-sum-nums[i])] += count / 2
			current[abs(-sum+nums[i])] += count / 2
		}
		prev = current
		fmt.Println(current)
	}
	if _, ok := prev[S]; !ok {
		return 0
	}
	if S == 0 {
		return prev[S]
	}
	return prev[S] / 2
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	//fmt.Println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
	fmt.Println(findTargetSumWays([]int{1, 2, 1}, 0))
}
