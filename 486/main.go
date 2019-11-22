package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func PredictTheWinner1(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	win := make([][]int, n)
	for i := 0; i < n; i++ {
		win[i] = make([]int, n)
		win[i][i] = nums[i]
	}

	// win[i][j] = max(nums[i]-win[i+1][j], nums[j]-win[i][j-1])
	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			win[i][j] = max(nums[i]-win[i+1][j], nums[j]-win[i][j-1])
		}
	}
	fmt.Println(win[0][n-1])
	return win[0][n-1] > 0
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	win := make([]int, n)
	for i := 0; i < n; i++ {
		win[i] = nums[i]
	}

	// win[i][j] = max(nums[i]-win[i+1][j], nums[j]-win[i][j-1])
	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			win[i] = max(nums[i]-win[i+1], nums[j]-win[i])
		}
		fmt.Println(nums[:n-l])
	}
	fmt.Println(win[0])
	return win[0] > 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(PredictTheWinner1([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner1([]int{1, 5, 233, 7}))
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}
