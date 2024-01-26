package main

import "fmt"

func checkRecord(s string) bool {
	countA := 0
	countL := 0
	for i := range s {
		c := s[i]
		if c == 'A' {
			countL = 0 // stupid
			countA++
			if countA >= 2 {
				return false
			}
		} else if c == 'L' {
			countL++
			if countL >= 3 {
				return false
			}
		} else {
			countL = 0
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
