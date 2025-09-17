package main

import "fmt"

// f(n, k) = f()
func combine2(n int, k int) [][]int {
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

type retType = [][]int

func combine1(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}

	dp := make([][]retType, k+1)
	// 1st line
	dp[0] = make([]retType, n+1)
	for j := 0; j <= n; j++ {
		dp[0][j] = retType{{}}
	}

	for i := 1; i <= k; i++ {
		dp[i] = make([]retType, n+1)
		dp[i][i-1] = retType{}

		for j := i; j <= n; j++ {
			left := make(retType, len(dp[i][j-1]))
			copy(left, dp[i][j-1])
			upLeft := make(retType, len(dp[i-1][j-1]))
			//copy(upLeft, dp[i-1][j-1])
			for t, com := range dp[i-1][j-1] {
				tmp := make([]int, len(com))
				copy(tmp, com)
				upLeft[t] = append(tmp, j)
			}
			dp[i][j] = make(retType, len(left)+len(upLeft))
			copy(dp[i][j], left)
			copy(dp[i][j][len(left):], upLeft)
		}
	}
	return dp[k][n]
}

func combine(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}

	// 1st line
	prevLine := make([]retType, n+1)
	for j := 0; j <= n; j++ {
		prevLine[j] = retType{{}}
	}
	for i := 1; i <= k; i++ {
		crtLine := make([]retType, n+1)
		crtLine[i-1] = retType{}

		for j := i; j <= n-k+i; j++ {
			crtLine[j] = [][]int{}
			left := crtLine[j-1]
			for _, elm := range left {
				tmp := make([]int, len(elm))
				copy(tmp, elm)
				crtLine[j] = append(crtLine[j], tmp)
			}
			upLeft := prevLine[j-1]
			for _, elm := range upLeft {
				tmp := make([]int, len(elm))
				copy(tmp, elm)
				tmp = append(tmp, j)
				crtLine[j] = append(crtLine[j], tmp)
			}
		}
		prevLine = crtLine
	}
	return prevLine[n]
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(combine(4, 2))
	//fmt.Println(combine(3, 3))
	//fmt.Println(combine(1, 1))
	//fmt.Println(combine(1, 0))
	fmt.Println(combine(5, 4))
	//fmt.Println(combine1(5, 4))
}
