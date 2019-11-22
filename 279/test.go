package main

import "fmt"

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	var count int
	seen := make([]bool, n)
	q := []int{n}
	for len(q) > 0 {
		count++
		var newQ []int
		for k := 0; k < len(q); k++ {
			x := q[k]
			for i := 1; i*i <= x; i++ {
				if x == i*i {
					return count
				}
				next := x - i*i
				if !seen[next] {
					newQ = append(newQ, next)
					seen[next] = true
				}
			}
		}
		q = newQ
	}
	return count
}
