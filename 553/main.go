package main

import "fmt"

func optimalDivision(nums []int) string {
	if len(nums) == 0 {
		return ""
	} else if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	} else if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	ret := fmt.Sprintf("%d/(%d", nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		ret += fmt.Sprintf("/%d", nums[i])
	}
	ret += ")"
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(optimalDivision([]int{1000, 100, 10, 2}))
}
