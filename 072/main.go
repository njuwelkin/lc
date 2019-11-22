package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			tmp := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				tmp++
			}
			dp[i][j] = min(dp[i][j], tmp)
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
