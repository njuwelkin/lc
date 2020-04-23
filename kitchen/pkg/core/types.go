package core

import (
	"time"
)

type OrderRequest struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Temp      string  `json:"temp"`
	ShelfLife int     `json:"shelfLife"`
	DecayRate float64 `json:"decayRate"`
}

type OrderTemp int

const (
	Hot OrderTemp = iota
	Cold
	Frozen
	InvalidTemp
)

type OrderStatus int

const (
	Accepted OrderStatus = iota
	Cooking
	OnShelf
	Dispatched
	Discarded
)

type Order struct {
	ID        string
	Name      string
	Temp      OrderTemp
	ShelfLife int

	RemainLefe int
	StartTime  time.Time
	UpdateTime time.Time
	Status     OrderStatus
}

func (o *Order) Value() float64 {
	return float64(o.RemainLefe) / float64(o.ShelfLife)
}

/*
type Kitchen interface {
	PlaceOrder(*OrderRequest)
}

type ShelveMgr interface {
	Put(*Order)
	Pick(*Order) error
}
*/

// abstract interface of cook, shelfPutter, shelfPicker, courier
type Colleague interface {
	Notify(*Order)
	AddSuccessor(Colleague)
}

type BaseColleague struct {
	colleagues []Colleague
}

func NewBaseColleague() *BaseColleague {
	return &BaseColleague{colleagues: []Colleague{}}
}

func (bc *BaseColleague) Send(order *Order) {
	for _, colleague := range bc.colleagues {
		colleague.Notify(order)
	}
}

func (bc *BaseColleague) AddSuccessor(colleague Colleague) {
	bc.colleagues = append(bc.colleagues, colleague)
}
