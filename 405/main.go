package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	unum := uint32(num)
	dict := "0123456789abcdef"
	ret := make([]byte, 8)
	idx := len(ret) - 1
	for unum != 0 {
		tmp := unum & 15
		//fmt.Printf("%c", dict[tmp])
		ret[idx] = dict[tmp]
		idx--
		unum >>= 4
	}
	return string(ret[idx+1:])
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(toHex(3))
	fmt.Println(toHex(-3))
}
