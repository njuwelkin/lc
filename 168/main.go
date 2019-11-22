package main

import "fmt"

func convertToTitle(n int) string {
	buf := []byte{}
	for n != 0 {
		n--
		buf = append(buf, 'A'+byte(n%26))
		n /= 26
	}
	for i, j := 0, len(buf)-1; i < j; {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}
