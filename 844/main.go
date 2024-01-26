package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	counts, countt := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		if i >= 0 && s[i] == '#' {
			fmt.Println("s, get #")
			i--
			counts++
			continue
		}
		if j >= 0 && t[j] == '#' {
			fmt.Println("t, get #")
			j--
			countt++
			continue
		}
		if i >= 0 && counts > 0 {
			fmt.Println("s, backtrace")
			i--
			counts--
			continue
		}
		if j >= 0 && countt > 0 {
			fmt.Println("t, backtrace")
			j--
			countt--
			continue
		}
		if (i < 0 || j < 0) || s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("bxj##tw", "bxj###tw"))
}
