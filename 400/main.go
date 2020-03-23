package main

import "fmt"

// 1-9 9
// 10-99 90*2=180
// 100-999 900*3=270
// 9*1 90*2 900*3 ...
func findNthDigit(n int) int {
	base := 1
	var i, length int
	for i, length = 1, 1; n > 9*i*length; i, length = i*10, length+1 {
		n -= 9 * i * length
		base *= 10
	}
	num := base + (n-1)/length
	fmt.Println(num)
	offset := (n - 1) % length
	for i = 0; i < length-offset-1; i++ {
		num /= 10
	}
	return num % 10
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(findNthDigit(1))
	//fmt.Println(findNthDigit(9))
	//fmt.Println(findNthDigit(10))
	//fmt.Println(findNthDigit(11))
	fmt.Println(findNthDigit(187))
	fmt.Println(findNthDigit(188))
	fmt.Println(findNthDigit(189))
	fmt.Println(findNthDigit(190))
}
