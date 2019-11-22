package main

import (
	"fmt"
	"math"
)

func find132pattern1(nums []int) bool {
	var i, j, k int
	prev1 := math.MaxInt64

	n := len(nums)

	for i = 0; i < n-2; i++ {
		if nums[i] > nums[i+1] || nums[i] >= prev1 {
			continue
		}
		prev1 = nums[i]

		prev2 := math.MaxInt64
		for k = n - 1; k > i+1; k-- {
			if nums[k] <= nums[i] || nums[k] >= prev2 {
				continue
			}
			if k-1 > i+1 && nums[k-1] > nums[i] && nums[k] > nums[k-1] {
				continue
			}
			prev2 = nums[k]

			for j = i + 1; j < k; j++ {
				if nums[j] > nums[i] && nums[j] > nums[k] {
					return true
				}
			}
		}
	}
	return false
}

func find132pattern2(nums []int) bool {
	n := len(nums)
	max := make([][]int, n)
	for i := 0; i < n; i++ {
		max[i] = make([]int, n)
		max[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			max[i][j] = max[i][j-1]
			if nums[j] > max[i][j] {
				max[i][j] = nums[j]
			}
		}
	}

	prev1 := math.MaxInt64
	for i := 0; i < n-2; i++ {
		if nums[i] > nums[i+1] || nums[i] >= prev1 {
			continue
		}
		prev1 = nums[i]

		for k := n - 1; k > i+1; k-- {
			if nums[k] > nums[i] && max[i+1][k-1] > nums[k] {
				return true
			}
		}
	}
	return false
}

func find132pattern3(nums []int) bool {
	n := len(nums)
	min := make([]int, n)
	minNum := math.MaxInt64
	for i, num := range nums {
		if num < minNum {
			minNum = num
		}
		min[i] = minNum
	}

	for k := n - 1; k >= 2; k-- {
		for j := k - 1; j >= 1; j-- {
			if nums[j] > nums[k] && min[j-1] < nums[k] {
				return true
			}
		}
	}
	return false
}

func find132pattern(nums []int) bool {
	n := len(nums)
	jStack := []int{}
	kVal := math.MinInt64
	for i := n - 1; i >= 0; i-- {
		if nums[i] < kVal {
			return true
		}
		for len(jStack) > 0 && nums[i] > jStack[len(jStack)-1] {
			kVal = jStack[len(jStack)-1]
			jStack = jStack[:len(jStack)-1]
		}
		jStack = append(jStack, nums[i])
	}
	return false
}

func main() {
	//fmt.Println("vim-go")
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
	fmt.Println(find132pattern([]int{-2, 1, 2, -2, 1, 2}))
	fmt.Println(find132pattern([]int{8, 10, 4, 6, 5}))
}
