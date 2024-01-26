package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
	h, w   int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	h, w := len(matrix), len(matrix[0])
	for i := 1; i < w; i++ {
		matrix[0][i] += matrix[0][i-1]
	}
	for j := 1; j < h; j++ {
		matrix[j][0] += matrix[j-1][0]
	}
	for j := 1; j < h; j++ {
		for i := 1; i < w; i++ {
			matrix[j][i] += matrix[j-1][i] + matrix[j][i-1] - matrix[j-1][i-1]
		}
	}
	return NumMatrix{
		matrix: matrix,
		h:      len(matrix),
		w:      len(matrix[0]),
	}
}

func (this *NumMatrix) get(row, col int) int {
	if row < 0 || col < 0 {
		return 0
	}
	return this.matrix[row][col]
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.get(row2, col2) - this.get(row2, col1-1) - this.get(row1-1, col2) + this.get(row1-1, col1-1)
}

func main() {
	fmt.Println("vim-go")
	obj := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	fmt.Println(obj.matrix)
	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
	fmt.Println(obj.SumRegion(1, 1, 1, 1))
	fmt.Println(obj.SumRegion(4, 3, 4, 3))
}
