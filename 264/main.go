package main

import "fmt"

func nthUglyNumber(n int) int {
	buf := make([]int, n)
	buf[0] = 1
	fact := [3]int{2, 3, 5}
	idx := [3]int{0, 0, 0}
	for i := 1; i < n; i++ {
		minVal := 1 << 62
		var minIdx int
		for k := 0; k < 3; k++ {
			tmp := fact[k] * buf[idx[k]]
			if tmp < minVal {
				minVal = tmp
				minIdx = k
			} else if tmp == minVal {
				idx[k]++
			}
		}
		buf[i] = minVal
		idx[minIdx]++
	}
	fmt.Println(buf)
	return buf[n-1]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(nthUglyNumber(10))
}
