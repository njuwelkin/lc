package main

import (
	"fmt"
	"math"
)

type SummaryRanges struct {
	Intervals [][]int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	ret := SummaryRanges{
		Intervals: [][]int{{math.MinInt64, math.MinInt64}, {math.MaxInt64, math.MaxInt64}},
	}
	return ret
}

func search(intv [][]int, val int) int {
	var i, j int
	for i, j = 0, len(intv); i < j; {
		mid := (i + j) / 2
		if intv[mid][0] < val {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func (this *SummaryRanges) AddNum(val int) {
	idx := search(this.Intervals, val)
	if this.Intervals[idx-1][1] >= val || this.Intervals[idx][0] == val {
		return
	}
	if this.Intervals[idx-1][1] < val-1 && this.Intervals[idx][0] > val+1 {
		this.Intervals = append(this.Intervals, make([]int, 2))
		for i := len(this.Intervals) - 1; i > idx; i-- {
			this.Intervals[i] = this.Intervals[i-1]
		}
		this.Intervals[idx] = []int{val, val}
		return
	}
	if val == this.Intervals[idx-1][1]+1 {
		this.Intervals[idx-1][1] = val
	}
	if val == this.Intervals[idx][0]-1 {
		this.Intervals[idx][0] = val
	}
	if this.Intervals[idx-1][1] == this.Intervals[idx][0] {
		this.Intervals[idx-1][1] = this.Intervals[idx][1]
		for i := idx; i < len(this.Intervals)-1; i++ {
			this.Intervals[i] = this.Intervals[i+1]
		}
		this.Intervals = this.Intervals[:len(this.Intervals)-1]
	}

}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.Intervals[1 : len(this.Intervals)-1]
}

func main() {
	fmt.Println("vim-go")
	obj := Constructor()
	obj.AddNum(1)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(3)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(7)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(2)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(6)
	fmt.Println(obj.GetIntervals())
}
