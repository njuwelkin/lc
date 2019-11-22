package ds

type Stack []int

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	if len(*s) != 0 {
		ret := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return ret
	}
	return Null
}

func (s Stack) Top() int {
	if len(s) != 0 {
		return s[len(s)-1]
	}
	return Null
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
