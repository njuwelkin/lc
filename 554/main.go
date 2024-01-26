package main

import "fmt"

func leastBricks(wall [][]int) int {
	edges := map[int]int{}
	for i := range wall {
		for j := 1; j < len(wall[i]); j++ {
			edges[wall[i][j-1]]++
			wall[i][j] += wall[i][j-1]
		}
	}
	max := 0
	for _, v := range edges {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(leastBricks([][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}))
}
