package main

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	for i := 0; i < len(A)-2; i++ {
		for j := i + 1; j < len(A)-1; j++ {

		}
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}
