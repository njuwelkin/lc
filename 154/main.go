package main

import "fmt"

func findMin(nums []int) int {
	for i, j := 0, len(nums)-1; ; {
		if i == j || nums[i] < nums[j] {
			return nums[i]
		}
		mid := (i + j + 1) / 2
		if nums[i] == nums[j] {
			if nums[mid] == nums[j] {
				if mid == j {
					return nums[i]
				} else {
					left := findMin(nums[i:mid])
					if left < nums[mid] {
						return left
					}
					return findMin(nums[mid+1 : j+1])
				}
			} else if nums[mid] < nums[j] {
				j = mid
				i++
			} else {
				i = mid + 1
			}
		} else {
			if nums[mid] <= nums[j] {
				j = mid
				i++
			} else {
				i = mid + 1
			}
		}
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{1, 3, 5}))
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
	fmt.Println(findMin([]int{2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2}))
}
