package main

import (
	"fmt"
	"sort"
)

func find(window []int, val int) int {
	i, j := 0, len(window)
	for i < j {
		mid := (i + j) / 2
		if window[mid] == val {
			return mid
		} else if window[mid] < val {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func move(window []int, vDel, vIns int) {
	if vDel == vIns {
		return
	}
	idxDel := find(window, vDel)
	idxIns := find(window, vIns)
	if idxDel >= idxIns {
		for i := idxDel; i > idxIns; i-- {
			window[i] = window[i-1]
		}
		window[idxIns] = vIns
	} else {
		for i := idxDel; i < idxIns-1; i++ {
			window[i] = window[i+1]
		}
		window[idxIns-1] = vIns
	}
}

func midian(window []int) float64 {
	mid := len(window) / 2
	if len(window)%2 == 0 {
		return float64(window[mid]+window[mid-1]) / 2
	} else {
		return float64(window[mid])
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	if k > len(nums) {
		return nil
	}
	ret := []float64{}
	window := make([]int, k)
	copy(window, nums[:k])
	sort.Ints(window)
	fmt.Println(window)
	ret = append(ret, midian(window))
	for i := k; i < len(nums); i++ {
		move(window, nums[i-k], nums[i])
		ret = append(ret, midian(window))
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 4, 2, 3}, 4))
}
