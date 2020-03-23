package main

import "fmt"

func sqrt(area int) int {
	i, j := 0, area+1
	for i < j {
		mid := (i + j) / 2
		res := mid * mid
		if res == area {
			return mid
		} else if res < area {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i - 1
}

func constructRectangle2(area int) []int {
	w := sqrt(area)
	l := w
	for {
		res := w * l
		if res == area {
			return []int{w, l}
		} else if res < area {
			w++
		} else {
			l--
			w = area / l
		}
	}
}

func constructRectangle(area int) []int {
	for l := sqrt(area); l > 0; l-- {
		if area%l == 0 {
			return []int{area / l, l}
		}
	}
	return nil
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(sqrt(2))
	fmt.Println(constructRectangle(1))
	fmt.Println(constructRectangle(2))
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(5))
	fmt.Println(constructRectangle(6))
	fmt.Println(constructRectangle(7))
	fmt.Println(constructRectangle(8))
	fmt.Println(constructRectangle(100))
}
