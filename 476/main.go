package main

import "fmt"

func findComplement(num int) int {
	var mask int
	for mask = 1; mask < num; mask <<= 1 {
	}
	mask--
	return mask & ^num
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}
