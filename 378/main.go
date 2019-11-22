package main

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if k > n*n {
		return -1
	}

	l, u := matrix[0][0], matrix[n-1][n-1]+1
	for l < u {
		mid := l + (u-l)/2

		count := 0
		for i, j := 0, n-1; i < n && j >= 0; {
			if matrix[i][j] <= mid {
				i++
				count += j + 1
				if count > k {
					break
				}
			} else {
				j--
			}
		}
		fmt.Println("mid, count:", mid, count)
		if count < k {
			l = mid + 1
		} else {
			u = mid
		}
	}
	return l
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8))
	fmt.Println(kthSmallest([][]int{{-5}}, 1))
}
