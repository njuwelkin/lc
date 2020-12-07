package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func puzzle() (string, int) {
	a, b := rand.Intn(20), rand.Intn(10)
	return fmt.Sprintf("%d + %d = ", a, b), a + b
}

func checkArgs() int {
	if len(os.Args) == 1 {
		return 20
	}
	ret, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("wrong args")
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	start := time.Now()
	rand.Seed(start.Unix())
	countWrong := 0
	total := checkArgs()
	for i := 0; i < total; i++ {
		q, a := puzzle()
		//for j := 0; j < 2; j++ {
		for {
			fmt.Printf("第%d题: %s", i+1, q)
			var rawinput string
			fmt.Scanln(&rawinput)
			input, err := strconv.Atoi(rawinput)
			if err == nil && input == a {
				fmt.Println("正确")
				break
			}
			countWrong++
			fmt.Println("错")
		}
		//}
	}
	d := time.Since(start)
	fmt.Printf("总共%d题，错误%d次，用时%f分钟\n", total, countWrong, d.Minutes())
}
