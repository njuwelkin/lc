package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	for l, r := 0, m*n; l < r; {
		mid := (l + r) / 2
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		}
		if val < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 13))
}
