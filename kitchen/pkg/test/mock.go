package test

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

// mock shelf
type shelf struct {
	content map[string]*core.Order
}

func (s *shelf) Put(o *core.Order) {
	s.content[o.ID] = o
}

func (s *shelf) Pick(o *core.Order) error {
	_, found := s.content[o.ID]
	if !found {
		return core.ResourceNotFound
	}
	delete(s.content, o.ID)
	return nil
}

// mock kitchen
type mockKitchen struct {
	s     *shelf
	Count int
}

func NewKitchen() *mockKitchen {
	return &mockKitchen{
		s: &shelf{map[string]*core.Order{}},
	}
}

func (k *mockKitchen) Send(o *core.Order, e core.Event) {
	if e != core.Discarded {
		panic("invalid event")
	}
	k.Count++
}

func (k *mockKitchen) GetShelf() core.Shelf {
	return k.s
}

func (k *mockKitchen) Find(o *core.Order) bool {
	_, found := k.s.content[o.ID]
	return found
}
