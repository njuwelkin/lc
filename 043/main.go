package main

import "fmt"

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 == 0 || l2 == 0 {
		return ""
	}
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}

	retArray := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		var val1 int = int(num1[i] - '0')
		power1 := l1 - 1 - i
		for j := l2 - 1; j >= 0; j-- {
			var val2 int = int(num2[j] - '0')
			power2 := l2 - 1 - j
			power := power1 + power2
			retArray[power] += val1 * val2
		}
	}
	//fmt.Println(retArray[:retArrayLen])

	carry := 0
	for i := 0; i < len(retArray); i++ {
		retArray[i] += carry
		carry = retArray[i] / 10
		retArray[i] %= 10
	}
	if retArray[len(retArray)-1] == 0 {
		retArray = retArray[:len(retArray)-1]
	}

	byteArray := make([]byte, len(retArray))
	for i := 0; i < len(retArray); i++ {
		byteArray[i] = '0' + byte(retArray[len(retArray)-1-i])
	}
	return string(byteArray)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
}
