package main

import "fmt"

func div2(b []int) ([]int, bool) {
	rem := 0
	for i := 0; i < len(b); i++ {
		tmp := b[i] % 2
		b[i] = (rem*10 + b[i]) / 2
		rem = tmp
	}
	if b[0] == 0 {
		return b[1:], rem != 0
	}
	return b, rem != 0
}
func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 0
	}
	a %= 1337

	var f func([]int) int
	f = func(b []int) int {
		fmt.Println(b)
		if len(b) == 1 && b[0] == 1 {
			return a
		}
		b, rem := div2(b)
		ret := f(b)
		if rem {
			return (a * ret * ret) % 1337
		} else {
			return (ret * ret) % 1337
		}
	}
	return f(b)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(superPow(2, []int{}))
	fmt.Println(superPow(2, []int{3}))
	fmt.Println(superPow(2, []int{1, 0}))
}
