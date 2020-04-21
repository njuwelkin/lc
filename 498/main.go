package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	ret := make([]int, 0, m*n)
	for line := 0; line < m+n-1; line++ {
		if line%2 == 1 {
			for j := min(n-1, line); j >= 0; j-- {
				i := line - j
				if i >= m {
					break
				}
				ret = append(ret, matrix[i][j])
			}
		} else {
			for i := min(m-1, line); i >= 0; i-- {
				j := line - i
				if j >= n {
					break
				}
				ret = append(ret, matrix[i][j])
			}
		}
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		//{7, 8, 9},
	}))
}
