package main

import "fmt"

func isPerfectSquare(num int) bool {
	for i, j := 0, num+1; i < j; {
		mid := (i + j) / 2
		square := mid * mid
		if square == num {
			fmt.Println(mid)
			return true
		} else if square < num {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isPerfectSquare(9))
	fmt.Println(isPerfectSquare(8))
	fmt.Println(isPerfectSquare(7))
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(0))
}
