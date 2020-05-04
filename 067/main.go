package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ret := make([]byte, len(a)+1)
	var i int
	var carry byte = 0
	for i = 1; i <= len(b); i++ {
		ret[len(ret)-i] = a[len(a)-i] + b[len(b)-i] - '0' + carry
		if ret[len(ret)-i] > '1' {
			carry = 1
			ret[len(ret)-i] -= 2
		} else {
			carry = 0
		}
	}

	for ; i <= len(a); i++ {
		ret[len(ret)-i] = a[len(a)-i] + carry
		if ret[len(ret)-i] > '1' {
			carry = 1
			ret[len(ret)-i] -= 2
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		ret[0] = '1'
		return string(ret)
	} else {
		return string(ret[1:])
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
