package main

import (
	"fmt"
)

func distance(a []int, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}

func numberOfBoomerangs(points [][]int) int {
	ret := 0
	for i := 0; i < len(points); i++ {
		dm := map[int]int{}
		for j := 0; j < len(points); j++ {
			dist := distance(points[i], points[j])
			if n, ok := dm[dist]; !ok {
				dm[dist] = 1
			} else {
				ret += 2 * n
				dm[dist]++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}))
}
