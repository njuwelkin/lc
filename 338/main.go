package main

import "fmt"

func countBits(num int) []int {
	ret := make([]int, num+1)
	ret[0] = 0
	for i := 1; ; i *= 2 {
		for j := i; j < i*2; j++ {
			if j > num {
				return ret
			}
			ret[j] = 1 + ret[j-i]
		}
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(countBits(10))
}
