package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	minDelta := 1 << 62
	for i := 0; i < len(nums)-2; i++ {
		tar := target - nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			tmp := tar - (nums[j] + nums[k])
			if tmp == 0 {
				return target
			}
			if abs(tmp) < abs(minDelta) {
				minDelta = tmp
			}
			if tmp > 0 {
				j++
			} else {
				k--
			}
		}
	}
	return target - minDelta
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
