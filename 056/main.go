package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	var i, j int
	for i, j = 0, 1; j < len(intervals); j++ {
		if intervals[i][1] >= intervals[j][0] {
			if intervals[j][1] > intervals[i][1] {
				intervals[i][1] = intervals[j][1]
			}
		} else {
			i++
			intervals[i] = intervals[j]
		}
	}
	i++
	return intervals[:i]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
