package main

import "fmt"

func fingerPrint(s string) uint32 {
	var ret uint32 = 0
	for _, c := range s {
		ret |= 1 << (c - 'a')
	}
	return ret
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	f1, f2 := fingerPrint(s1), fingerPrint(s2[:len(s1)])
	if f1&f2 == f1 {
		return true
	}
	fmt.Printf("%s: %b\n", s1, f1)

	for i := len(s1); i < len(s2); i++ {
		f2 ^= 1 << (s2[i-len(s1)] - 'a')
		f2 |= 1 << (s2[i] - 'a')
		fmt.Printf("%s: %b\n", s2[i-len(s1)+1:i+1], f2)
		if f1&f2 == f1 {
			fmt.Println(s2[i-len(s1)+1 : i+1])
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("adc", "dcda"))
}
