package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if j, found := m[num]; found {
			if i-j <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	} else if t == 0 {
		return containsNearbyDuplicate(nums, k)
	}

	m := map[int]int{}
	check := func(num int) bool {
		if _, found := m[num/t]; found {
			return true
		}
		if v, found := m[num/t+1]; found && v-num <= t {
			return true
		}
		if v, found := m[num/t-1]; found && num-v <= t {
			return true
		}
		return false
	}

	for i := 0; i < k && i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num -= t - 1
		}
		if check(num) {
			fmt.Println(m)
			return true
		}
		m[num/t] = num
	}
	for i := k; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num -= t - 1
		}
		if check(num) {
			return true
		}
		m[num/t] = num
		j := i - k
		delete(m, nums[j]/t)
	}
	return false
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	} else if t == 0 {
		return containsNearbyDuplicate(nums, k)
	}
	m := map[int][]int{}
	for i, num := range nums {
		if num < 0 {
			num -= t - 1
		}
		if v, found := m[num/t]; found && i-v[0] <= k {
			return true
		}
		if v, found := m[num/t+1]; found && i-v[0] <= k && v[1]-num <= t {
			return true
		}
		if v, found := m[num/t-1]; found && i-v[0] <= k && num-v[1] <= t {
			return true
		}
		m[num/t] = []int{i, num}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{-1, -1}, 1, -1))
	fmt.Println(containsNearbyAlmostDuplicate([]int{-3, 3}, 2, 4))
}
