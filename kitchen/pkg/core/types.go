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
	Accept Event = iota
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
)

type Order struct {
	// static info parsed from order request
	ID        string
	Name      string
	Temp      OrderTemp
	ShelfLife float64
	DecayRate float64

	// dynamic info
	RemainLife float64
	UpdateTime time.Time
	Status     string

	// courier reports a estimate picking time after receive a picking request
	//   this will enable kitchen system to estimate the remain value when picking
	// system will prefer to discard a order with least picking value when shelf is full
	EstimatePickTime time.Time

	// once courier
	IsCanceled chan struct{}
	// courier will block on this channel when order is not ready for picking(not cooked, etc)
	// in this case the block will not happen as food can always been cooked immediately
	//   while courier always need to sleep some seconds. but still have it here for functional integrity.
	IsOnShelf chan struct{}
}

func (o *Order) Value() float64 {
	return o.RemainLife / float64(o.ShelfLife)
}

func (o *Order) UpdateRemainLife(now time.Time, onOverFlow bool) {
	age := now.Unix() - o.UpdateTime.Unix()
	shelfDecayModifier := 1
	if onOverFlow {
		shelfDecayModifier = 2
	}
	o.UpdateTime = now
	o.RemainLife -= o.DecayRate * float64(age) * float64(shelfDecayModifier)
}

func (o *Order) EstimatePickValue(onOverFlow bool) float64 {
	age := o.EstimatePickTime.Unix() - o.UpdateTime.Unix()
	shelfDecayModifier := 1
	if onOverFlow {
		shelfDecayModifier = 2
	}
	remainLife := o.RemainLife - o.DecayRate*float64(age)*float64(shelfDecayModifier)
	return remainLife / float64(o.ShelfLife)
}

type Kitchen interface {
	Send(*Order, Event)
	GetShelf() Shelf
}

type Shelf interface {
	Put(*Order)
	Pick(id string) error
	SetKitchen(Kitchen)
}

// abstract interface of cook, courier
type Colleague interface {
	Notify(*Order)
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
