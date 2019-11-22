package main

import "fmt"

func maxRotateFunction(A []int) int {
	sum := 0
	ret := 0
	for i, v := range A {
		sum += v
		ret += i * v
	}

	n := len(A)
	tmp := ret
	for i := n - 1; i >= 0; i-- {
		tmp = tmp + sum - n*A[i]
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
}
