package main

import "fmt"

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	if n < 4 {
		return 2
	}

	n /= 2
	if n%2 != 0 {
		return 4 * lastRemaining(n/2)
	} else {
		return 4*lastRemaining(n/2) - 2
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(lastRemaining(4))
}
