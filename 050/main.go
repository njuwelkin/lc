package main

import "fmt"

func myPow(x float64, n int) float64 {
	isNeg := false
	if n < 0 {
		isNeg = true
		n = -n
	}

	var ret float64 = 1.0
	for n > 0 {
		if n&1 != 0 {
			ret *= x
		}
		x *= x
		n >>= 1
	}
	if isNeg {
		return 1.0 / ret
	} else {
		return ret
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(myPow(2.0, 3))
	fmt.Println(myPow(3.0, 3))
	fmt.Println(myPow(2.0, -3))
}
