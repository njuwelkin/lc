package main

import "fmt"

func rand7() int {
	return 1
}

func rand10() int {
    for {
    	a := rand7() - 1
	b := rand7() - 1
	idx := a * 7 + b
	if idx >= 40 {
		continue
	}
	return idx %10 + 1
    }
}

func main() {
	fmt.Println("vim-go")
}
