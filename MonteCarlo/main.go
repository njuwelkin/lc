package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var _once sync.Once

type CoinChangeState struct {
	amount int
	coins  []int
}

func NewCCState(amount int, coins []int) CoinChangeState {
	return CoinChangeState{
		amount: amount,
		coins:  coins,
	}
}

func (ccs *CoinChangeState) getUntriedActions() []int {
	ret := []int{}
	for _, c := range ccs.coins {
		if c <= ccs.amount {
			ret = append(ret, c)
		}
	}
	return ret
}

func (ccs *CoinChangeState) performAction(action int) CoinChangeState {
	return CoinChangeState{
		amount: ccs.amount - action,
		coins:  ccs.coins,
	}
}
func (ccs *CoinChangeState) simulate() int {
	fmt.Println(ccs.amount)
	if ccs.amount == 0 {
		return 0
	}
	return rand.Intn(ccs.amount)
}

type node struct {
	state    CoinChangeState
	parent   *node
	children []*node
	visits   int
	value    int
}

func newNode(state CoinChangeState, parent *node) *node {
	return &node{
		state:    state,
		parent:   parent,
		children: []*node{},
	}
}

func _select(p *node) *node {
	if len(p.children) == 0 {
		return p
	}

	exploration_factor := 1.0
	var selectedChild *node = nil
	maxWeight := 0.0
	for _, child := range p.children {
		weight := (float64(child.value) / (float64(child.visits) + 0.000001)) +
			exploration_factor*math.Sqrt(math.Log(float64(p.visits+1))/(float64(child.visits)+0.000001))
		if weight > maxWeight {
			maxWeight = weight
			selectedChild = child
		}
	}
	return _select(selectedChild)
}

func expand(p *node) *node {
	untriedActions := p.state.getUntriedActions()
	if len(untriedActions) == 0 {
		return nil
	}
	i := rand.Intn(len(untriedActions))
	action := untriedActions[i]
	new_state := p.state.performAction(action)
	new_child := newNode(new_state, p)
	p.children = append(p.children, new_child)
	return new_child
}

func simulate(p *node) int {
	return p.state.simulate()
}

func backpropagate(p *node, value int) {
	for ; p != nil; p = p.parent {
		p.visits++
		p.value += value
	}
}

type MctsTree = *node

func NewMctsTree(initState CoinChangeState) MctsTree {
	return newNode(initState, nil)
}

func (mt MctsTree) Run(max_iterations int) CoinChangeState {
	root := mt
	for i := 0; i < max_iterations; i++ {
		p := _select(root)
		new_child := expand(p)
		if new_child != nil {
			value := simulate(new_child)
			if value == 0 {
				break
			}
			backpropagate(new_child, value)
		}
	}

	maxVisits := 0
	var ret *node = nil
	for _, child := range root.children {
		if child.visits > maxVisits {
			ret = child
		}
	}
	return ret.state
}

func main() {
	fmt.Println("vim-go")
	_once.Do(func() {
		rand.Seed(time.Now().Unix())
	})

	initState := NewCCState(63, []int{1, 5, 10, 25})
	fmt.Println(NewMctsTree(initState).Run(1000))
}
