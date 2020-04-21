package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	ret := make([]byte, len(num2)+1)
	fmt.Println(len(ret))

	var carry byte = 0
	var i int
	for i = 0; i < len(num1); i++ {
		ret[len(ret)-i-1] = carry + num1[len(num1)-i-1] + num2[len(num2)-i-1] - '0'
		carry = 0
		if ret[len(ret)-i-1] > '0'+9 {
			ret[len(ret)-i-1] -= 10
			carry = 1
		}
	}
	for ; i < len(num2); i++ {
		ret[len(ret)-i-1] = carry + num2[len(num2)-i-1]
		carry = 0
		if ret[len(ret)-i-1] > '0'+9 {
			ret[len(ret)-i-1] -= 10
			carry = 1
		}
	}
	if carry == 0 {
		return string(ret[1:])
	} else {
		ret[0] = '1'
		return string(ret)
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(addStrings("123", "456"))
	fmt.Println(addStrings("123", "789"))
	fmt.Println(addStrings("323", "789"))
	fmt.Println(addStrings("1", "9"))
}
