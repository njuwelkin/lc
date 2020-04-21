package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func totalNQueens(n int) int {
	ret := 0

	rows := make([]int, n)
	var f func(line int)
	f = func(line int) {
		if line == n {
			ret++
			return
		}
		for i := 0; i < n; i++ {
			conflict := false
			for j := 0; j < line; j++ {
				if rows[j] == i || abs(i-rows[j]) == line-j {
					conflict = true
					break
				}
			}
			if conflict {
				continue
			}
			rows[line] = i
			f(line + 1)
		}
	}
	for i := 0; i < n/2; i++ {
		rows[0] = i
		f(1)
	}
	ret *= 2
	if n%2 != 0 {
		rows[0] = n / 2
		f(1)
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(totalNQueens(1))
	fmt.Println(totalNQueens(2))
	fmt.Println(totalNQueens(3))
	fmt.Println(totalNQueens(4))
	fmt.Println(totalNQueens(5))
	fmt.Println(totalNQueens(6))
}
