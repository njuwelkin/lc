package main

import "fmt"

func canSplit(nums []int, m int, sum int, maxSum int) bool {
	i := 0
	for i < len(nums) {
		if sum > maxSum*m {
			return false
		}
		subSum := 0
		for ; i < len(nums); i++ {
			subSum += nums[i]
			if subSum > maxSum {
				sum -= subSum - nums[i]
				break
			}
		}
		m--
	}
	return true
}

func splitArray(nums []int, m int) int {
	sum := 0
	max := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	var i, j int
	for i, j = max, sum+1; i < j; {
		mid := (i + j) / 2
		if canSplit(nums, m, sum, mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(canSplit([]int{7, 2, 5, 10, 8}, 2, 32, 18))
	fmt.Println(canSplit([]int{7, 2, 5, 10, 8}, 2, 32, 19))
}
