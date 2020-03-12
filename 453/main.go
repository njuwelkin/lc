package main

import "fmt"

// (n-1) * m + sum = (min + m) * n
// m = sum - n * min
func minMoves(nums []int) int {
	fmt.Println(nums)
	n := len(nums)
	sum := 0
	min := 1 << 62
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}

	return sum - n*min
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 2, 2, 3}))
	fmt.Println(minMoves([]int{2, 2, 2, 3}))
	fmt.Println(minMoves([]int{3, 3, 3}))
	fmt.Println(minMoves([]int{1, 1, 1}))
}
