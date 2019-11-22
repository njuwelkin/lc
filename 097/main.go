package main

import "fmt"

func isInterleave1(s1 string, s2 string, s3 string) bool {
	if s1 == "" {
		return s2 == s3
	} else if s2 == "" {
		return s1 == s2
	}

	m := len(s1)
	n := len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s3[i-1] == s1[i-1]
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s3[j-1] == s2[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (s1[i-1] == s3[i+j-1]) && dp[i-1][j] ||
				(s2[j-1] == s3[i+j-1]) && dp[i][j-1]
		}
	}
	return dp[m][n]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := map[[3]int]bool{}

	var isInterleaveRec func(int, int, int) bool
	isInterleaveRec = func(i1 int, i2 int, i3 int) bool {
		if ret, ok := dp[[3]int{i1, i2, i3}]; ok {
			return ret
		}
		if i1 == l1 {
			dp[[3]int{i1, i2, i3}] = s2[i2:] == s3[i3:]
			return dp[[3]int{i1, i2, i3}]
		}
		if i2 == l2 {
			dp[[3]int{i1, i2, i3}] = s1[i1:] == s3[i3:]
			return dp[[3]int{i1, i2, i3}]
		}
		dp[[3]int{i1, i2, i3}] = (s1[i1] == s3[i3]) && isInterleaveRec(i1+1, i2, i3+1) ||
			(s2[i2] == s3[i3]) && isInterleaveRec(i1, i2+1, i3+1)
		return dp[[3]int{i1, i2, i3}]
	}

	return isInterleaveRec(0, 0, 0)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}
