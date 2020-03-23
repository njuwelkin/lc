package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Solution struct {
	r    float64
	x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	a := rand.Float64()
	l := rand.Float64()
	return []float64{this.x + this.r*l*math.Sin(a), this.y + this.r*l*math.Cos(a)}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor(1, 0, 0)
	fmt.Println(obj.RandPoint())
	fmt.Println(obj.RandPoint())
	fmt.Println(obj.RandPoint())
}
