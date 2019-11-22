package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	ret := len(nums) + 1
	var i, j int
	for i, j = 0, 0; j < len(nums); {
		fmt.Println(i, j, nums[i:j], sum)
		if sum < s {
			sum += nums[j]
			j++
		} else {
			sum -= nums[i]
			fmt.Println(nums[i:j])
			if j-i < ret {
				ret = j - i
			}
			i++
		}
	}
	for sum >= s {
		fmt.Println(i, j, nums[i:j], sum)
		sum -= nums[i]
		i++
	}
	if j-i+1 < ret {
		ret = j - i + 1
	}
	if ret > len(nums) {
		return 0
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
