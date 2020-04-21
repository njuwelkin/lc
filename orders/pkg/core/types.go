package core

import (
	"time"
)

type OrderRequest struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Temp      string `json:"temp"`
	ShelfLife int    `json:"shelfLife"`
	DecayRate float  `json:"decayRate"`
}

type OrderTemp int

const (
	Hot OrderTemp = iota
	Cold
	Frozen
	InvalidTemp
)

var tempNames = map[string]OrderTemp{
	"hot":    Hot,
	"cold":   Cold,
	"frozen": Frozen,
}

type Order struct {
	ID         string
	Name       string
	Temp       OrderTemp
	ShelfLife  int
	RemainLefe int
	UpdateTime time.Time
}

func NewOrder(req *OrderRequest) *order {

}

func (o *Order) Value() float {
	return float(o.RemainLefe) / o.ShelfLife
}

type Kitchen interface {
	PlaceOrder(*OrderRequest)
}

type ShelveMgr interface {
	Put(*Order)
	Pick(*Order) error
}

type CookerMgr interface {
	Notify(*Order)
}

type CourierMgr interface {
	Notify(*Order)
}
