package main

import "fmt"

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	} else if x < 0 || x%10 == 0 {
		return false
	}
	y := 0
	for y < x {
		t := x % 10
		x /= 10
		y = y*10 + t
	}
	if y == x || y/10 == x {
		return true
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(10))
}
