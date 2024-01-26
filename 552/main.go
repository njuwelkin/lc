package main

import "fmt"

func checkRecord1(n int) int {
	/*
		ret := 1               // 0a0l
		ret += n               // 1a0l
		ret += n               // 0a1l
		ret += n * (n - 1)     // 1a1l
		ret += n * (n - 1) / 2 // 0a2l
		if n > 2 {
			ret += n * (n - 1) * (n - 2) / 2 // 1a2l
		}
		return ret
	*/

	if n <= 2 {
		return (3*n*n+n)/2 + 1
	}
	//return (n*n*n+3*n)/2 + 1

	/*
	 * y1 = k1x + b1, y2 = k2x + b2
	 * 1. (y1 + y2) % x = ((k1+k2)x + b1+b2)%x = (b1+b2)%x = (y1%x + y2%x) % x
	 * 2. (y1 * y2) % x = (f(k,x,b)*x + b1*b2)%x = (b1*b2)%x = (y1%x * y2%x) % x
	 */
	x := 1000000007
	b := n % x // n = kx + b, so n^2 = f(k,x,b)*x + b^2
	return (((((b*b)%x)*b)%x+3*b)/2 + 1) % x
}

func checkRecord(n int) int {
	x := 1000000007
	dp := [2][3]int{{1, 1, 0}, {1, 0, 0}} // dp[i][j]: i = countA; j == countTailL
	for i := 2; i <= n; i++ {
		tmp := [2][3]int{
			{dp[0][0], dp[0][1], dp[0][2]},
			{dp[1][0], dp[1][1], dp[1][2]},
		}
		// p
		dp[0][0] = (tmp[0][0] + tmp[0][1] + tmp[0][2]) % x
		dp[1][0] = tmp[1][0] + tmp[1][1] + tmp[1][2]
		// a
		dp[1][0] += tmp[0][0] + tmp[0][1] + tmp[0][2]
		dp[1][0] %= x
		// l
		dp[0][1] = tmp[0][0] % x
		dp[0][2] = tmp[0][1] % x
		dp[1][1] = tmp[1][0] % x
		dp[1][2] = tmp[1][1] % x

	}
	ret := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			ret += dp[i][j]
		}
	}
	return ret % x
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(checkRecord(1))
	fmt.Println(checkRecord(2))
	fmt.Println(checkRecord(3))
	fmt.Println(checkRecord(4))
	fmt.Println(checkRecord(5))
}
