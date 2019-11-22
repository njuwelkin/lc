package main

import "fmt"

func test(n int) {
	buf := make([]int, n)
	for i := 1; i <= n; i++ {
		for j := 0; j < n; j++ {
			if (j+1)%i == 0 {
				buf[j] = 1 - buf[j]
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		count += buf[i]
	}
	fmt.Println(n, buf, count)
}

func bulbSwitch(n int) int {
	return 0
}

func main() {
	fmt.Println("vim-go")
	for i := 1; i < 30; i++ {
		test(i)
	}
}
