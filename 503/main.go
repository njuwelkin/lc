package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	ret := make([]int, len(nums))

	maxIdx := 0
	maxNum := -1 << 63
	for i, num := range nums {
		if num > maxNum {
			maxIdx = i
			maxNum = num
		}
		ret[i] = -1
	}

	stack := []int{}
	for j := 1; j <= len(nums); {
		i := (maxIdx + j) % len(nums)
		if len(stack) == 0 || nums[stack[len(stack)-1]] >= nums[i] {
			stack = append(stack, i)
			j++
		} else {
			topIdx := stack[len(stack)-1]
			ret[topIdx] = nums[i]
			stack = stack[:len(stack)-1]
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
