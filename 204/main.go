package main

import "fmt"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	notPrime := make([]bool, n)
	count := 2 // 0 and 1 are have as not prime
	for i := 3; i < n; i++ {
		for j := 2; j < i; j++ {
			if !notPrime[j] && i%j == 0 {
				count++
				notPrime[i] = true
				break
			}
		}
	}
	return n - count
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(4))
	fmt.Println(countPrimes(5))
	//fmt.Println(countPrimes(10))
}
