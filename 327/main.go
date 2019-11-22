package main

import "fmt"

func merge(left, right, sum []int, lower, upper int) int {
	ret := 0
	idxLower, idxUpper := 0, 0
	for i := 0; i < len(left); i++ {
		val := left[i] + lower
		for ; idxLower < len(right) && right[idxLower] < val; idxLower++ {
		}
		val = left[i] + upper + 1
		for ; idxUpper < len(right) && right[idxUpper] < val; idxUpper++ {
		}
		ret += idxUpper - idxLower
	}
	tmp := make([]int, len(left))
	copy(tmp, left)
	left = tmp

	idx := 0
	var i, j int
	for i, j = 0, 0; i < len(left) && j < len(right); idx++ {
		if left[i] < right[j] {
			sum[idx] = left[i]
			i++
		} else {
			sum[idx] = right[j]
			j++
		}
	}
	for ; i < len(left); i++ {
		sum[idx] = left[i]
		idx++
	}
	return ret
}

func countRangeSumInternal(sum []int, lower int, upper int) int {
	if len(sum) <= 1 {
		return 0
	}
	mid := len(sum) / 2
	ret := countRangeSumInternal(sum[:mid], lower, upper)
	ret += countRangeSumInternal(sum[mid:], lower, upper)
	ret += merge(sum[:mid], sum[mid:], sum, lower, upper)
	fmt.Println(sum, ret)
	return ret
}

func countRangeSum(nums []int, lower int, upper int) int {
	sum := make([]int, len(nums)+1)
	// sum [i, j) = sum[j]-sum[i]; sum [i, j] = sum[j+1] - sum[i]
	// sum [0, n] = sum[n+1] - sum[0]
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	fmt.Println(sum)
	return countRangeSumInternal(sum, lower, upper)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
