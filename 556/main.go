package main

import "fmt"

func nextGreaterElement(n int) int {
	if n <= 0 || n >= 1<<32 {
		return -1
	}
	power1 := 1

	var tmp int
	var n1, n2 int
	prev := 0
	for tmp = n; tmp > 0; tmp /= 10 {
		n1 = tmp % 10
		if n1 < prev {
			break
		}
		prev = n1
		power1 *= 10
	}
	if tmp == 0 {
		return -1
	}

	power2 := 1
	for tmp = n; tmp > 0; tmp /= 10 {
		n2 = tmp % 10
		if n2 > n1 {
			break
		}
		power2 *= 10
	}

	ret := n - n1*power1 + n2*power1 - n2*power2 + n1*power2

	reverse := 0
	for p1 := ret % power1; p1 > 0; p1 /= 10 {
		d := p1 % 10
		reverse = 10*reverse + d
	}
	ret -= ret % power1
	ret += reverse
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(nextGreaterElement(1239871)) // 1271389
}
