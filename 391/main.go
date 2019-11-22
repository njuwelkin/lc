package main

import (
	"fmt"
	"math"
)

type Point struct {
	i, j int
}

func area(rect []int) int {
	return (rect[2] - rect[0]) * (rect[3] - rect[1])
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

func isRectangleCover(rectangles [][]int) bool {
	m := map[Point]int{}
	totalArea := 0
	top, left := math.MaxInt64, math.MaxInt64
	bottom, right := 0, 0
	for _, rect := range rectangles {
		totalArea += area(rect)
		for _, idx := range [][]int{{0, 1}, {0, 3}, {2, 1}, {2, 3}} {
			p := Point{rect[idx[0]], rect[idx[1]]}
			if _, ok := m[p]; !ok {
				m[p] = 0
			}
			m[p]++
		}
		top = min(top, rect[0])
		left = min(left, rect[1])
		bottom = max(bottom, rect[2])
		right = max(right, rect[3])
	}
	if (right-left)*(bottom-top) != totalArea {
		return false
	}
	for _, p := range []Point{{top, left}, {top, right}, {bottom, left}, {bottom, right}} {
		if m[p] != 1 {
			return false
		}
		m[p] = 0
	}
	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}))

	fmt.Println(isRectangleCover([][]int{
		{1, 1, 2, 3},
		{1, 3, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
	}))

	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{3, 2, 4, 4},
	}))

	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4},
	}))
}
