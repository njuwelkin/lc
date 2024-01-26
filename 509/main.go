package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	fn_2 := 0
	fn_1 := 1
	fn := 1
	for i := 2; i <= n; i++ {
		fn = fn_2 + fn_1
		fn_2, fn_1 = fn_1, fn
	}
	return fn
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
}
