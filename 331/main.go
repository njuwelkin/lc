package main

import "fmt"

func isValidSerialization(preorder string) bool {
	array := strings.Split(preorder, ",")
	if len(array) == 0 {
		return true
	}

	i := 0
	var f func() bool
	f = func() bool {
		if i >= len(array) {
			return false
		}
		if array[i] == "#" {
			i++
			return true
		}
		i++
		return f() && f()
	}
	return f() && i == len(array)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("1,#"))
	fmt.Println(isValidSerialization("9,#,#,1"))
	fmt.Println(isValidSerialization("9,#,92,#,#"))
}
