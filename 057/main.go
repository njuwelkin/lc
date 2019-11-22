package main

import "fmt"

func searchLeft(intervals [][]int, left int) int {
	var i, j int
	for i, j = 0, len(intervals); i < j; {
		mid := (i + j) / 2
		if intervals[mid][1] < left {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func searchRight(intervals [][]int, right int) int {
	var i, j int
	for i, j = 0, len(intervals); i < j; {
		mid := (i + j) / 2
		if intervals[mid][0] <= right {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func insert(intervals [][]int, newInterval []int) [][]int {
	left := searchLeft(intervals, newInterval[0])
	right := searchRight(intervals, newInterval[1])
	fmt.Println(left, right)

	if left == right {
		//append
		intervals = append(intervals, []int{0, 0})
		for i := len(intervals) - 1; i > right; i-- {
			intervals[i] = intervals[i-1]
		}
		intervals[left] = newInterval
	} else {
		intervals[left][0] = min(intervals[left][0], newInterval[0])
		intervals[left][1] = max(intervals[right-1][1], newInterval[1])
		if right-left != 1 {
			var i, j int
			for i, j = left+1, right; j < len(intervals); {
				intervals[i] = intervals[j]
				i++
				j++
			}
			intervals = intervals[:i]
		}
	}
	return intervals
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 9}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {13, 16}}, []int{11, 12}))
}
