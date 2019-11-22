package main

import "fmt"

/*
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
*/

func flood1(grid [][]int, i, j int) int {
	hit := 0
	q := [][]int{{i, j}}
	grid[i][j] = -1
	for len(q) != 0 {
		head := q[0]
		for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := []int{head[0] + d[0], head[1] + d[1]}
			if next[0] < 0 || next[0] >= len(grid) ||
				next[1] < 0 || next[1] >= len(grid[0]) ||
				grid[next[0]][next[1]] == -1 {
				continue
			}
			if grid[next[0]][next[1]] == 1 {
				hit++
				continue
			}
			grid[next[0]][next[1]] = -1
			q = append(q, next)
		}
		q = q[1:]
	}
	return hit
}

func islandPerimeter1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ret := 0
	for j := 0; j < len(grid[0]); j++ {
		for _, i := range []int{0, len(grid) - 1} {
			if grid[i][j] == 0 {
				ret += flood1(grid, i, j)
			} else if grid[i][j] == 1 {
				ret++
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for _, j := range []int{0, len(grid[0]) - 1} {
			if grid[i][j] == 0 {
				ret += flood1(grid, i, j)
			} else if grid[i][j] == 1 {
				ret++
			}
		}
	}
	return ret
}

func flood2(grid [][]int, i, j int) int {
	fmt.Println("flood", i, j)
	hit := 0
	q := [][]int{{i, j}}
	grid[i][j] = -1
	for len(q) != 0 {
		head := q[0]
		for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := []int{head[0] + d[0], head[1] + d[1]}
			if next[0] < 0 || next[0] >= len(grid) ||
				next[1] < 0 || next[1] >= len(grid[0]) ||
				grid[next[0]][next[1]] == 0 {
				hit++
				continue
			}
			if grid[next[0]][next[1]] == -1 {
				continue
			}
			grid[next[0]][next[1]] = -1
			q = append(q, next)
		}
		q = q[1:]
	}
	return hit
}

func islandPerimete2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return flood2(grid, i, j)
			}
		}
	}
	return 0
}

func islandPerimeter(grid [][]int) int {
	ret := 0
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			ret++
		}
		if grid[m-1][j] == 1 {
			ret++
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != grid[i-1][j] {
				ret++
			}
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			ret++
		}
		if grid[i][n-1] == 1 {
			ret++
		}
	}
	for j := 1; j < n; j++ {
		for i := 0; i < m; i++ {
			if grid[i][j] != grid[i][j-1] {
				ret++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0}}))
	fmt.Println(islandPerimeter([][]int{{1, 0}}))
	fmt.Println(islandPerimeter([][]int{{0, 1}}))
}
