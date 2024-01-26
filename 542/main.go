package main

import "fmt"

func updateMatrix2(mat [][]int) [][]int {
	queue := [][]int{}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j, 0}) // enqueue
			} else {
				mat[i][j] = 10000 // not visited
			}
		}
	}
	for len(queue) > 0 {
		i, j, v := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:] // delqueue
		if i > 0 && mat[i-1][j] == 10000 {
			mat[i-1][j] = v + 1
			queue = append(queue, []int{i - 1, j, v + 1})
		}
		if i < len(mat)-1 && mat[i+1][j] == 10000 {
			mat[i+1][j] = v + 1
			queue = append(queue, []int{i + 1, j, v + 1})
		}
		if j > 0 && mat[i][j-1] == 10000 {
			mat[i][j-1] = v + 1
			queue = append(queue, []int{i, j - 1, v + 1})
		}
		if j < len(mat[i])-1 && mat[i][j+1] == 10000 {
			mat[i][j+1] = v + 1
			queue = append(queue, []int{i, j + 1, v + 1})
		}
	}
	return mat
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func updateMatrix(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				mat[i][j] = 10000
				if i > 0 {
					mat[i][j] = min(mat[i][j], mat[i-1][j]+1)
				}
				if j > 0 {
					mat[i][j] = min(mat[i][j], mat[i][j-1]+1)
				}
			}
		}
	}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[i]) - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if i != len(mat)-1 {
					mat[i][j] = min(mat[i][j], mat[i+1][j]+1)
				}
				if j != len(mat[i])-1 {
					mat[i][j] = min(mat[i][j], mat[i][j+1]+1)
				}
			}
		}
	}
	return mat
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}
