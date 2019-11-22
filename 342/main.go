package main

import "fmt"

func isPowerOfFour(num int) bool {
	return num != 0 && (num&(num-1) == 0) && (num == 1 || num%10 == 6 || num%10 == 4)
}

func main() {
	fmt.Println("vim-go")
}
