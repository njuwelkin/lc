package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)

	var getDepth func(i, j int) int
	getDepth = func(i, j int) int {
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		ret := 0
		for _, idx := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			x, y := i+idx[0], j+idx[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if matrix[x][y] <= matrix[i][j] {
				continue
			}
			ret = max(ret, getDepth(x, y))
		}
		dp[i][j] = ret + 1
		return dp[i][j]
	}

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret = max(ret, getDepth(i, j))
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}))
	fmt.Println(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}))
}
