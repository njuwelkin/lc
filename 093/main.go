package main

import "fmt"

func restoreIpAddresses(s string) []string {
	ret := []string{}
	buf := make([]int, 4)
	var dfs func(int, string)
	dfs = func(depth int, str string) {
		if depth == 4 {
			ret = append(ret, fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3]))
			return
		}
		buf[depth] = 0
		maxLen := 3
		if str[0] == '0' {
			maxLen = 1
		}
		for i := 0; i < maxLen && i < len(str); i++ {
			buf[depth] = buf[depth]*10 + int(str[i]-'0')
			fmt.Println(depth, maxLen, buf, str)
			if buf[depth] > 255 {
				break
			}
			if len(str)-1-i < 4-depth-1 {
				break
			}
			if len(str)-1-i > (4-depth-1)*3 {
				continue
			}
			dfs(depth+1, str[i+1:])
		}
	}
	dfs(0, s)
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("010010"))
}
