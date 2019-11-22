package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])

	dungeon[m-1][n-1] = max(1, 1-dungeon[m-1][n-1])

	for i := m - 2; i >= 0; i-- {
		dungeon[i][n-1] = max(1, dungeon[i+1][n-1]-dungeon[i][n-1])
	}
	for j := n - 2; j >= 0; j-- {
		dungeon[m-1][j] = max(1, dungeon[m-1][j+1]-dungeon[m-1][j])
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dungeon[i][j] = min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j])
			dungeon[i][j] = max(1, dungeon[i][j])
		}
	}
	fmt.Println(dungeon)
	return dungeon[0][0]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}))
}
