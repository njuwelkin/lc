package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var ret uint32
	//for num != 0 {
	for i := 0; i < 32; i++ {
		ret <<= 1
		ret |= num & 1
		num >>= 1
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(reverseBits(1))
	//fmt.Println(reverseBits(2))
	//fmt.Println(reverseBits(3))
	fmt.Println(reverseBits(43261596))
}
