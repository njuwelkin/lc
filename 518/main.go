package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}

func change1(amount int, coins []int) int {
	dp := map[[2]int]int{}
	var f func(int, int) int
	f = func(amount int, start int) int {
		if amount == 0 {
			return 1
		} else if amount < 0 || start == len(coins) {
			return 0
		}
		if v, found := dp[[2]int{amount, start}]; found {
			return v
		}
		dp[[2]int{amount, start}] = f(amount-coins[start], start) + f(amount, start+1)
		return dp[[2]int{amount, start}]
	}
	return f(amount, 0)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
}
