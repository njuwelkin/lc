package main

import (
	"fmt"
	"sort"
)

func minSessions(tasks []int, sessionTime int) int {
	sort.Ints(task)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minSessions([]int{1, 2, 3}, 3))
	fmt.Println(minSessions([]int{3, 1, 3, 1, 1}, 8))
	fmt.Println(minSessions([]int{1, 2, 3, 4, 5}, 15))
	fmt.Println(minSessions([]int{2, 2, 3, 3, 3, 5}, 9))
	fmt.Println(minSessions([]int{5, 5, 5, 5, 5, 5}, 9))
}
