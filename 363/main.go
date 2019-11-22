package main

import (
	"fmt"
	"math"
)

func getSum(matrix [][]int, top, left, bottom, right int) int {
	if top == 0 && left == 0 {
		return matrix[bottom][right]
	} else if top == 0 {
		return matrix[bottom][right] - matrix[bottom][left-1]
	} else if left == 0 {
		return matrix[bottom][right] - matrix[top-1][right]
	} else {
		return matrix[bottom][right] - matrix[bottom][left-1] - matrix[top-1][right] + matrix[top-1][left-1]
	}
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

	ret := math.MinInt64
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			for bottom := i; bottom < m; bottom++ {
				for right := j; right < n; right++ {
					sum := getSum(matrix, i, j, bottom, right)
					if sum == k {
						return k
					} else if sum < k && sum > ret {
						ret = sum
					}
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
}
