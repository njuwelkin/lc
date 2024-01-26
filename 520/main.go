package main

import "fmt"

func allLow(word string) bool {
	for i := range word {
		if !(word[i] >= 'a' && word[i] <= 'z') {
			return false
		}
	}
	return true
}

func allCap(word string) bool {
	for i := range word {
		if !(word[i] >= 'A' && word[i] <= 'Z') {
			return false
		}
	}
	return true
}

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return true
	}
	if word[0] >= 'A' && word[0] <= 'Z' {
		return allLow(word[1:]) || allCap(word[1:])
	} else if word[0] >= 'a' && word[0] <= 'z' {
		return allLow(word[1:])
	} else {
		return false
	}
}

func main() {
	fmt.Println("vim-go")
}
