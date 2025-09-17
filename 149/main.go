package main

import "fmt"

func normalize(faction [2]int) [2]int {
	if faction[0] == 0 {
		faction[1] = 1
		return faction
	}
	sign := 1
	if faction[0] < 0 {
		sign = -sign
		faction[0] = -faction[0]
	}
	if faction[1] < 0 {
		sign = -sign
		faction[1] = -faction[1]
	}
	x, y := faction[0], faction[1]
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		tmp := x % y
		x = y
		y = tmp
	}
	faction[0] /= x * sign
	faction[1] /= x
	return faction
}

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	m := map[[2][2]int][2]int{}
	for i := 0; i < len(points)-1; i++ {
		x0, y0 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x1, y1 := points[j][0], points[j][1]
			var k, b [2]int
			if x0 != x1 {
				k = normalize([2]int{y1 - y0, x1 - x0})
				b = normalize([2]int{(y1+y0)*(x1-x0) - (y1-y0)*(x1+x0), 2 * (x1 - x0)})
			} else {
				k = [2]int{1, 0}
				b = [2]int{1, x1}
			}
			key := [2][2]int{k, b}
			if v, ok := m[key]; !ok {
				m[key] = [2]int{2, i}
			} else {
				if v[1] == i {
					m[key] = [2]int{v[0] + 1, v[1]}
				}
			}
		}
	}
	ret := 0
	for _, v := range m {
		if v[0] > ret {
			ret = v[0]
		}
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}
