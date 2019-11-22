package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	n := len(height)
	ret := 0
	for i, j := 0, n-1; i < j; {
		if height[i] < height[j] {
			ret = max(ret, (j-i)*height[i])
			i++
		} else {
			ret = max(ret, (j-i)*height[j])
			j--
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
