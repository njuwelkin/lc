package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	countZero := 0
	productExceptZero := 1
	for _, num := range nums {
		if num != 0 {
			productExceptZero *= num
		} else {
			countZero++
			if countZero > 1 {
				break
			}
		}
	}
	if countZero > 1 {
		for i := range nums {
			nums[i] = 0
		}
	} else if countZero == 0 {
		for i, num := range nums {
			nums[i] = productExceptZero / num
		}
	} else {
		for i, num := range nums {
			if num == 0 {
				nums[i] = productExceptZero
			} else {
				nums[i] = 0
			}
		}
	}
	return nums
}

func main() {
	fmt.Println([]int{})
	fmt.Println()
}
