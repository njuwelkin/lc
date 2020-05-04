package main

import "fmt"

// f(n) = f(n-1) + f(n-2) = 2*f(n-2)+f(n-3) = 3*f(n-3)+2*f(n-4) = 5*f(n-4) + 3*f(n-5) 
//   == 
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	fn_1, fn_2 := 1, 1
	fn := 0
	for i := 2; i <= n; i++ {
		fn = fn_1 + fn_2
		fn_1, fn_2 = fn, fn_1
	}
	return fn
}

func main() {
	fmt.Println("vim-go")
}
