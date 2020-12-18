package main

import "fmt"

func isSelfCrossing(x []int) bool {
	var vertex [][2]int
	cx, cy := 0, 0
	for i := 0; i < len(x); i++ {
		vertex = append(vertex, [2]int{cx, cy})

		switch i % 4 {
		case 0:
			cy += x[i]
		case 1:
			cx -= x[i]
		case 2:
			cy -= x[i]
		case 3:
			cx += x[i]
		}
		if i >= 3 {
			fmt.Println(vertex)

		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4}))
	fmt.Println(isSelfCrossing([]int{1, 1, 1, 1}))
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4, 1, 2, 1, 1}))
}
