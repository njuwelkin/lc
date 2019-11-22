package main

import "fmt"

// f(k, n) = {{m}+f(k-1, n-m) : for m = 9 to 1}
// f(m, 1) = {{m}}
func combinationSum3(k int, n int) [][]int {
	var f func(int, int, int) [][]int
	f = func(k, n, minVal int) [][]int {
		if k == 1 {
			if n > minVal && n < 10 {

			} else {
				return nil
			}
		}
		for i := minVal + 1; i < 10; i++ {
			f(k-1, n-i, i)
		}
	}
}

func main() {
	fmt.Println("vim-go")
}
