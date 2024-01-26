package main

import "fmt"

func findRotateSteps(ring string, key string) int {
	n := len(ring)
	pos := 0
	ret := 0
	rotate := func() {
		pos = (pos + 1) % n
		ret++
	}

	push := func() {
		ret++
	}

	for i := 0; i < len(key); i++ {
		for ring[pos] != key[i] {
			rotate()
		}
		push()
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findRotateSteps("godding", "gd"))
	fmt.Println(findRotateSteps("godding", "godding"))
}
