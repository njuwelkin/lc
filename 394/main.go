package main

import (
	"fmt"
	"strconv"
)

/*
E ::= null | TE
T ::= str | num[E]
*/
func isNumber(s string) bool {
	return s[0] >= '0' && s[0] <= '9'
}

func decodeString(s string) string {
	i := 0
	w := ""
	getw := func() {
		if i >= len(s) {
			w = ""
			return
		}
		if s[i] == '[' {
			i++
			w = "["
		} else if s[i] == ']' {
			i++
			w = "]"
		} else if s[i] >= '0' && s[i] <= '9' {
			start := i
			for i = i + 1; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			}
			w = s[start:i]
		} else {
			start := i
			for i = i + 1; i < len(s) && !(s[i] == '[' || s[i] == ']' || s[i] >= '0' && s[i] <= '9'); i++ {
			}
			w = s[start:i]
		}
	}

	getw()

	var E func() string
	var T func() string
	E = func() string {
		if w == "" || w == "]" {
			return ""
		}
		return T() + E()
	}
	T = func() string {
		if isNumber(w) {
			repeat, err := strconv.Atoi(w)
			if err != nil {
				panic("")
			}
			getw()
			if w != "[" {
				panic("")
			}
			getw()
			content := E()
			if w != "]" {
				panic("")
			}
			getw()

			ret := ""
			for ; repeat > 0; repeat-- {
				ret += content
			}
			return ret
		} else {
			ret := w
			getw()
			return ret
		}
	}
	return E()
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(decodeString("3[ab]2[cd3[ef]]gh"))
}
