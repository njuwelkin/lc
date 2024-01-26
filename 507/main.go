package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == 2*num
}

func main() {
	fmt.Println("vim-go")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d: %t\n", i, checkPerfectNumber(i))
	}
}
