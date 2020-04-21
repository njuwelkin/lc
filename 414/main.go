package main

import "fmt"

func thirdMax(nums []int) int {
	if len(nums) < 3 {
		return 1 << 62
	}
	bottom := len(nums) - 1
	top := 0
	for i := 0; i < 3; i++ {
		for ; bottom > i && nums[bottom] <= nums[bottom-1]; bottom-- {
		}
		if bottom == i {
			break
		}
		for j := bottom; j > i; j-- {
			if nums[j] > nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				top = j
			}
		}
		if top > 2 {
			break
		}
	}
	fmt.Println(nums)
	return nums[2]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(thirdMax([]int{1, 2, 3, 4, 5, 6, 7}))
}
