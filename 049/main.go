package main

import "fmt"

func fingerPrint(s string) [26]int {
	var ret [26]int
	for i := 0; i < len(s); i++ {
		ret[s[i]-'a']++
	}
	return ret
}

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		fp := fingerPrint(str)
		if _, found := m[fp]; !found {
			m[fp] = []string{str}
		} else {
			m[fp] = append(m[fp], str)
		}
	}

	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
