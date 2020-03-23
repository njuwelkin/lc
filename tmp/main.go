package main

import (
	"fmt"
)

func isPalindrome(n int) bool {
	if n%10 == 0 {
		//return false
	}
	m := 0
	for m < n {
		tmp := n % 10
		m = 10*m + tmp
		n /= 10
	}
	if m > n {
		m /= 10
	}
	return m == n
}

func main() {
	fmt.Println("vim-go")

	bound := 10
	for i := 1; i < 10; i++ {
		max := 0
		maxM, maxN := 1, 1
		for m := bound - 1; m >= bound/10; m-- {
			if m*m < max {
				break
			}
			for n := m; n >= bound/10; n-- {
				product := m * n
				if product < max {
					break
				}
				if isPalindrome(product) {
					max = product
					maxM, maxN = m, n
				}
			}
		}
		fmt.Printf("%d*%d=%d, mod 1337=%d\n", maxM, maxN, max, max%1337)
		bound *= 10
	}

}
