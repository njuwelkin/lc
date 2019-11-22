package main

import "fmt"

func getPermutation(n int, k int) string {
	candidate := make([]byte, n)
	for i := 0; i < n; i++ {
		candidate[i] = '0' + byte(i+1)
	}

	fact := 1
	for i := 1; i < n; i++ {
		fact *= i
	}
	k--
	for i := 0; i < n-1; i++ {
		fmt.Println(fact, k, string(candidate))
		idx := i + k/fact
		if idx != i {
			tmp := candidate[idx]
			for j := idx; j > i; j-- {
				candidate[j] = candidate[j-1]
			}
			candidate[i] = tmp
		}
		k %= fact
		fact /= (n - i - 1)
	}
	return string(candidate)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(getPermutation(3, 3)) //213
	fmt.Println(getPermutation(4, 9)) //2314
	fmt.Println(getPermutation(3, 2)) //132
}
