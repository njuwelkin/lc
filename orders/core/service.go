package core

import (
	"time"
)

type Temperature string

const (
	Hot    Temperature = "hot"
	Cold               = "cold"
	Frozen             = "frozen"
)

type RawOrder struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Temp      Temperature `json:"temp"`
	ShelfLife int         `json:"shelfLife"`
	DecayRate float       `json:"decayRate"`
}

type Order struct {
	*RawOrder
	StartTime  time.Time
	RemainLefe int
}

func (o *Order) Value() float {
	return float(o.RemainLefe) / o.ShelfLife
}

type Kitchen interface {
	PlaceOrder(*Order)
	CookDone(*Order)
}

type ShelveMgr interface {
	Put(*Order)
	Pick(*Order)
}

type CookerMgr interface {
	Notify(*Order)
}

type CourierMgr interface {
	Notify(*Order)
}
