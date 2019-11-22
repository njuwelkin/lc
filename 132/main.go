package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	cut := make([]int, n)
	for i := range cut {
		cut[i] = i
	}
	if s[1] == s[0] {
		cut[1] = 0
	}

	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if i == j || s[i] == s[j] && ((j == i+1) || dp[i+1][j-1]) {
				dp[i][j] = true
				if i == 0 {
					cut[j] = 0
				} else {
					cut[j] = min(cut[j], cut[i-1]+1)
				}
			}
		}
	}
	fmt.Println(cut)

	return cut[n-1]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minCut("aab"))
	fmt.Println(minCut("abcd"))
	fmt.Println(minCut("abad"))
	fmt.Println(minCut("aacddd"))
	fmt.Println(minCut("cabababcbc"))
}
