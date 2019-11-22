package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	ret := [][]int{}
	for i := 0; i < n-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			t := target - nums[i] - nums[j]
			for k, l := j+1, n-1; k < l; {
				sum := nums[k] + nums[l]
				if sum == t {
					if k == j+1 || nums[k] != nums[k-1] ||
						l == n-1 || nums[l] != nums[l+1] {
						ret = append(ret, []int{nums[i], nums[j], nums[k], nums[l]})
					}
					k++
					l--
				} else if sum < t {
					k++
				} else {
					l--
				}
			}
		}
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(fourSum([]int{-1, 0, -5, -2, -2, -4, 0, 1, -2}, -9))
}
