package test

import (
	"kitchen/pkg/core"
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
	ctx   *core.Context
	s     *shelf
	Count int
}

func NewKitchen(ctx *core.Context) *mockKitchen {
	return &mockKitchen{
		ctx: ctx,
		s:   &shelf{map[string]*core.Order{}},
	}
}

func (k *mockKitchen) Send(o *core.Order, e core.Event) {
	k.Count++
	k.ctx.Log.Infof("receive event %d for order %s", e, o.ID)
}

func (k *mockKitchen) GetShelf() core.Shelf {
	return k.s
}

func (k *mockKitchen) Find(o *core.Order) bool {
	k.ctx.Log.Infof("%v", k.s.content)
	_, found := k.s.content[o.ID]
	return found
}
