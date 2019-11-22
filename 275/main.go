package main

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	var i, j int
	for i, j = 0, n; i < j; {
		mid := (i + j) / 2
		if citations[mid] < n-mid {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return n - i
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{0, 1, 2, 5, 6}))
	fmt.Println(hIndex([]int{100}))
}
