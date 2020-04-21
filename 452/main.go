package main

import (
	"fmt"
	"sort"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] ||
			points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	ret := 1
	top := 0
	for i := 1; i < len(points); i++ {
		if points[i][0] <= points[top][1] {
			points[top][1] = min(points[top][1], points[i][1])
		} else {
			ret++
			points[top][1] = points[i][1]
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}
