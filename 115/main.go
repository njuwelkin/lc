package main

import "fmt"

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	if n == 0 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = 1
	}
	for i := 1; i <= n; i++ {
		for j := i; j <= m; j++ {
			dp[i][j] = dp[i][j-1]
			if s[j-1] == t[i-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
}
