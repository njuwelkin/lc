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
	Hot           OrderTemp     = iota // 0
	Cold                               // 1
	Frozen                             // 2
	InvalidTemp                        // 3
	CountTempType = InvalidTemp        // 3
)

type Event int

const (
	// accept a new order
	Accepted Event = iota
	// cook complete his job
	Cooked
	// food is relocated on shelf
	Moved
	// order is picked
	Picked
	// order is discarded
	Discarded
	// order is delivered
	Delivered
	// total number of event types
	CountEvent
)

type ShelfType int

const (
	SingleTempShelf ShelfType = iota
	OverflowShelf
)

type Order struct {
	// static info parsed from order request
	ID        string
	Name      string
	Temp      OrderTemp
	ShelfLife float64
	DecayRate float64

	// dynamic info
	RemainLife float64   // the remain shelf life, it reduce DecayRate*shelfDecayModifier every second
	UpdateTime time.Time // last time the RemainLife is updated
	Status     string    // human readable status for an order
	ShelfType  ShelfType // on singletemp shelf or overflow shelf

	// courier reports an estimate picking time after receive a picking request
	//   this will help kitchen system to estimate the remain value when picking
	// system will prefer to discard a order with minimal picking value when shelf is full
	EstimatePickTime time.Time

	// courier will block on this channel, if he has arrived at kitchen but food is not ready
	//   in this case cook can always complete immediately, the block will not happen,
	//   but still keep the function here
	Ready chan struct{}
	// courier will cancel the picking job once receiving a message from the cancel channel
	Cancel chan struct{}
}

/*
func (o *Order) Value() float64 {
	return o.RemainLife / float64(o.ShelfLife)
}
*/

func (o *Order) EstimatePickValue(onOverFlow bool) float64 {
	age := o.EstimatePickTime.Unix() - o.UpdateTime.Unix()
	shelfDecayModifier := 1
	if onOverFlow {
		shelfDecayModifier = 2
	}
	remainLife := o.RemainLife - o.DecayRate*float64(age)*float64(shelfDecayModifier)
	if remainLife < 0 {
		remainLife = 0
	}
	return remainLife / float64(o.ShelfLife)
}

type Kitchen interface {
	Send(*Order, Event)
	GetShelf() Shelf
}

type Shelf interface {
	Put(*Order)
	Pick(*Order) error
}

// abstract interface of cook, courier
type Colleague interface {
	Notify(*Order, Event)
	SetKitchen(Kitchen)
	GetOffWork()
}

type BaseColleague struct {
	Kitchen
}

func NewBaseColleague() *BaseColleague {
	return &BaseColleague{}
}

func (bc *BaseColleague) SetKitchen(k Kitchen) {
	bc.Kitchen = k
}
