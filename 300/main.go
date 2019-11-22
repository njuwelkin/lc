package main

import "fmt"

func search(path []int, num int) int {
	i, j := 0, len(path)
	for i < j {
		mid := (i + j) / 2
		if path[mid] < num {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := 1
	for i := 1; i < len(nums); i++ {
		idx := search(nums[:ret], nums[i])
		nums[idx] = nums[i]
		if idx == ret {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
