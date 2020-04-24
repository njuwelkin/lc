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
	CountTempType = InvalidTemp
)

type OrderStatus int

const (
	Accepted OrderStatus = iota
	Cooked
	OnShelf
	Picking
	Discarded
	Delivered
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

	WaitCook chan struct{}
}

func (o *Order) Value() float64 {
	return float64(o.RemainLefe) / float64(o.ShelfLife)
}

type Kitchen interface {
	Send(*Order)
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
	return BaseColleague{}
}

func (bc *BaseColleague) SetKitchen(k Kitchen) {
	bc.Kitchen = k
}

type job struct {
	f func()
}

func NewJob(f func()) *job {
	return &job{f: f}
}

func (j *job) Do() {
	j.f()
}
