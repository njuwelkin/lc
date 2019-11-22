package main

import (
	"fmt"
	"sort"
)

func nsum(n int, target int, nums []int) bool {
	if n == 1 {
		var i int
		for i = len(nums) - 1; i >= 0 && nums[i] == 0; i-- {
		}
		if nums[i] == target {
			nums[i] = 0
			return true
		}
		return false
	} else if n == 2 {
		for i, j := 0, len(nums)-1; i < j; {
			if nums[i] == 0 {
				i++
				continue
			} else if nums[j] == 0 {
				j--
				continue
			}
			sum := nums[i] + nums[j]
			if sum == target {
				nums[i], nums[j] = 0, 0
				return true
			} else if sum < target {
				i++
			} else {
				j--
			}
		}
		return false
	} else {
		for i := len(nums) - 1; i >= n-1; i-- {
			if nums[i] > target || nums[i] == 0 {
				continue
			}
			if nsum(n-1, target-nums[i], nums[:i]) {
				nums[i] = 0
				return true
			}
		}
		return false
	}
}

func makesquare(nums []int) bool {
	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	edge := sum / 4
	if sum == 0 || sum%4 != 0 || maxNum > edge {
		return false
	}

	sort.Ints(nums)

	countEdge := 4
	remain := len(nums)
	for i := 1; i*countEdge <= remain; i++ {
		for nsum(i, edge, nums) {
			remain -= i
			countEdge--
			if countEdge == 1 {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
}
