package main

import "fmt"

// f(n, k) = f()
func combine(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}

	ret := [][]int{}
	buf := make([]int, k)
	var f func(int, int)
	f = func(n, m int) {
		if m == 0 {
			tmp := make([]int, k)
			copy(tmp, buf)
			ret = append(ret, tmp)
			return
		}
		for i := n; i >= m; i-- {
			buf[m-1] = i
			f(i-1, m-1)
		}
	}
	f(n, k)

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(combine(4, 2))
	fmt.Println(combine(3, 3))
}
