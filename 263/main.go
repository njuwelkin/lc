package main

import "fmt"

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num%5 == 0 {
		num /= 5
	}
	for num%3 == 0 {
		num /= 3
	}
	return num&(num-1) == 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isUgly(0))
	fmt.Println(isUgly(2))
	fmt.Println(isUgly(3))
	fmt.Println(isUgly(5))
}
