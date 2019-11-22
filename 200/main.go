package main

import "fmt"

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	var flood func(int, int)
	flood = func(i, j int) {
		grid[i][j] = 0
		crtLevel := [][]int{{i, j}}
		for len(crtLevel) > 0 {
			nextLevel := [][]int{}
			for _, seed := range crtLevel {
				for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					x, y := seed[0]+d[0], seed[1]+d[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
						grid[x][y] = 0
						nextLevel = append(nextLevel, []int{x, y})
					}
				}
			}
			crtLevel = nextLevel
		}
	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				flood(i, j)
				ret++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
}
