package main

import "fmt"

func search(nums []int, target int) bool {
	for i, j := 0, len(nums); i < j; {
		mid := (i + j) / 2
		if nums[mid] == target || nums[i] == target || nums[j-1] == target {
			return true
		}

		if nums[i] < nums[j-1] {
			if target > nums[j-1] || target < nums[i] {
				return false
			}
			if nums[mid] < target {
				i = mid + 1
			} else {
				j = mid
			}
		} else if nums[i] == nums[j-1] {
			if nums[mid] == nums[i] {
				return search(nums[:mid], target) || search(nums[mid+1:], target)
			} else if nums[mid] > nums[i] {
				if target > nums[mid] || target < nums[i] {
					i = mid + 1
				} else {
					j = mid
				}
			} else {
				if target > nums[mid] && target < nums[i] {
					i = mid + 1
				} else {
					j = mid
				}
			}
		} else {
			if target > nums[j-1] && target < nums[i] {
				return false
			}
			if nums[mid] < nums[j-1] {
				if target > nums[mid] && target <= nums[j-1] {
					i = mid + 1
				} else {
					j = mid
				}
			} else if nums[mid] == nums[j-1] {
				j = mid
			} else if nums[mid] == nums[i] {
				i = mid + 1
			} else {
				if target > nums[mid] || target < nums[i] {
					i = mid + 1
				} else {
					j = mid
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{2, 0, 0, 1}, 1))
}
