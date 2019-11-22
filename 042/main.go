package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	leftBank, rightBank := height[0], height[n-1]

	ret := 0
	for i, j := 0, n-1; i < j; {
		if leftBank < rightBank {
			i++
			if height[i] < leftBank {
				ret += leftBank - height[i]
			} else {
				leftBank = height[i]
			}
		} else {
			j--
			if height[j] < rightBank {
				ret += rightBank - height[j]
			} else {
				rightBank = height[j]
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
