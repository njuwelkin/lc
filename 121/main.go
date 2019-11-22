package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ret := 0
	lowPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < lowPrice {
			lowPrice = prices[i]
		} else {
			profile := prices[i] - lowPrice
			if profile > ret {
				ret = profile
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
