package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getLine(matrix [][]int, top, bottom int) []int {
	ret := make([]int, len(matrix[0])+1)
	if top == 0 {
		copy(ret[1:], matrix[bottom])
	} else {
		for i := range matrix[bottom] {
			ret[i+1] = matrix[bottom][i] - matrix[top-1][i]
		}
	}
	return ret
}
func mergeSort(line []int, k int) int {
	if len(line) == 1 {
		return math.MaxInt64
	}
	mid := len(line) / 2
	ret := mergeSort(line[:mid], k)
	if ret == 0 {
		return 0
	}
	ret = min(ret, mergeSort(line[mid:], k))
	if ret == 0 {
		return 0
	}

	var i, j int
	for i, j = 0, mid; i < mid && j < len(line); {
		if line[j]-line[i] > k {
			i++
		} else {
			ret = min(ret, k-(line[j]-line[i]))
			if ret == 0 {
				return 0
			}
			j++
		}
	}

	tmp := make([]int, mid)
	copy(tmp, line[:mid])
	idx := 0
	for i, j = 0, mid; i < mid && j < len(line); {
		if tmp[i] < line[j] {
			line[idx] = tmp[i]
			i++
		} else {
			line[idx] = line[j]
			j++
		}
		idx++
	}
	for ; i < mid; i++ {
		line[idx] = tmp[i]
		idx++
	}
	return ret
}
func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	var i, j int
	for j = 1; j < n; j++ {
		matrix[0][j] += matrix[0][j-1]
	}
	for i = 1; i < m; i++ {
		matrix[i][0] += matrix[i-1][0]
	}

	for i = 1; i < m; i++ {
		for j = 1; j < n; j++ {
			matrix[i][j] += matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1]
		}
	}

	ret := math.MaxInt64
	for top := 0; top < m; top++ {
		for bottom := top; bottom < m; bottom++ {
			line := getLine(matrix, top, bottom)
			ret = min(ret, mergeSort(line, k))
		}
	}
	return k - ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 0))
}
