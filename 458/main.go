package main

import "fmt"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	k := minutesToTest / minutesToDie // how many times to test
	k++
	val := 1
	var i int
	for i = 0; val < buckets; i++ {
		val *= k
	}
	return i
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(poorPigs(1000, 15, 60))
	fmt.Println(poorPigs(15, 1, 2))
	fmt.Println(poorPigs(16, 1, 2))
	fmt.Println(poorPigs(17, 1, 2))
	fmt.Println(poorPigs(29, 1, 4))
	fmt.Println(poorPigs(30, 1, 4))
	fmt.Println(poorPigs(28, 1, 4))
}
