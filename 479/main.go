package main

import "fmt"

func power(n int) int {
	ret := 1
	fact := 10
	for n != 0 {
		if n&1 != 0 {
			ret *= fact
		}
		n >>= 1
		fact *= fact
	}
	return ret
}

func largestPalindrome(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 9
	} else if n%2 == 0 {
		a := power(n) - 1
		b := power(n) - power(n/2) + 1
		return (a % 1337) * (b % 1337) % 1337
	} else {
		a := power(n) - 5
		// 58(6)*(n-3)3
		b := 58*power(n-2) + (power(n-3)-1)*20/3 + 3
		return (a % 1337) * (b % 1337) % 1337
	}
}

func main() {
	fmt.Println("vim-go")
	for i := 1; i < 20; i++ {
		fmt.Println(largestPalindrome(i))
	}
}
