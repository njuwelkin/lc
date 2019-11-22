package main

import "fmt"

func numSquares1(n int) int {
	m := map[int]int{}
	sqrs := []int{}
	for i := 0; i < n; i++ {
		tmp := i * i
		if tmp > n {
			break
		}
		m[tmp] = i
		sqrs = append(sqrs, tmp)
	}
	fmt.Println(m)

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if _, found := m[i]; found {
			dp[i] = 1
			continue
		}
		dp[i] = i
		for _, k := range sqrs {
			if k > i-k || dp[i] == 2 {
				break
			}
			if 1+dp[i-k] < dp[i] {
				dp[i] = 1 + dp[i-k]
			}
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func numSquares2(n int) int {
	m := map[int]int{}
	sqrs := []int{}
	for i := 0; i < n; i++ {
		tmp := i * i
		if tmp > n {
			break
		}
		m[tmp] = i
		sqrs = append(sqrs, tmp)
	}
	fmt.Println(m)

	dp := make([]int, n+1)

	var f func(n int) int
	f = func(n int) int {
		if dp[n] != 0 {
			return dp[n]
		}
		if _, found := m[n]; found {
			dp[n] = 1
			return 1
		}
		dp[n] = n
		for i := 1; i < len(sqrs) && sqrs[i] <= n/2; i++ {
			if 1+f(n-sqrs[i]) < dp[n] {
				dp[n] = 1 + dp[n-sqrs[i]]
				if dp[n] == 2 {
					break
				}
			}
		}
		return dp[n]
	}
	f(n)
	fmt.Println(dp)
	return dp[n]
}

func numSquares3(n int) int {
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

func numSquares(n int) int {
	m := map[int]int{}
	for i := 0; i < n; i++ {
		tmp := i * i
		if tmp > n {
			break
		}
		m[tmp] = i
	}
	fmt.Println(m)

	var f func(n int, depth int) bool
	f = func(n int, depth int) bool {
		if _, found := m[n]; found {
			return true
		}
		if depth <= 1 {
			return false
		}
		for i := 1; i*i <= n/2; i++ {
			if f(n-i*i, depth-1) {
				return true
			}
		}
		return false
	}
	for i := 1; i < n; i++ {
		if f(n, i) {
			return i
		}
	}
	return n
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(10))
}
