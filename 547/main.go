package main

import "fmt"

func findCircleNum(M [][]int) int {
	n := len(M)
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(root int) {
		for i, v := range M[root] {
			if v != 0 && !visited[i] {
				visited[i] = true
				dfs(i)
			}
		}
	}

	ret := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			ret++
			visited[i] = true
			dfs(i)
		}
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
	fmt.Println(findCircleNum([][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}))
}
