package kitchen

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

type kitchen struct {
	// context
	ctx *core.Context

	// colleagues
	cookMgr    core.Colleague
	courierMgr core.Colleague

	shelf core.Shelf

	countDelieve int
	countDiscard int
}

func NewKitchen(ctx *core.Context, cookMgr, courierMgr core.Colleague, shelf core.Shelf) *kitchen {
	k := &kitchen{
		ctx:        ctx,
		cookMgr:    cookMgr,
		courierMgr: courierMgr,
		shelf:      shelf,
	}
	cookMgr.SetKitchen(k)
	courierMgr.SetKitchen(k)
	shelf.SetKitchen(k)
	return k
}

func (k *kitchen) PlaceOrder(req *core.OrderRequest) error {
	k.ctx.Log.WithField("ID", req.ID).Info("receive order")
	k.ctx.Log.Infof("receive order %+v", req)
	order, err := newOrder(req)
	if err != nil {
		k.ctx.Log.WithError(err).Warn("invalid order request")
		return err
	}
	k.ctx.Log.Infof("new order %+v", order)
	k.Send(order, core.Accept)
	return nil
}

func (k *kitchen) Send(order *core.Order, event core.Event) {
	go k.dispatch(order, event)
}

func (k *kitchen) Run() *kitchen {
	return k
}

func (k *kitchen) Stop() {
	k.cookMgr.GetOffWork()
	k.courierMgr.GetOffWork()
}

func (k *kitchen) GetShelf() core.Shelf {
	return k.shelf
}

func (k *kitchen) dispatch(order *core.Order, event core.Event) {
	switch event {
	case core.Accept:
		// notify cook manager to prepare the food
		k.cookMgr.Notify(order)
		// notify courier manager to dispatch a courier
		k.courierMgr.Notify(order)
	case core.Cooked:
		// food is ready, wakeup waiting courier
		order.IsOnShelf <- struct{}{}
		// notify cleaner to schedule a clean job
	case core.Moved:
		// notify cleaner to re-schedule the clean job
	case core.Discarded:
		k.countDiscard++
	case core.Delivered:
		k.countDelieve++
	}
}

var tempNames = map[string]core.OrderTemp{
	"hot":    core.Hot,
	"cold":   core.Cold,
	"frozen": core.Frozen,
}

func newOrder(req *core.OrderRequest) (*core.Order, error) {
	temp, found := tempNames[req.Temp]
	if !found {
		return nil, core.InvalidOrderRequest.WithField("Temp", req.Temp)
	}
	return &core.Order{
		ID:        req.ID,
		Name:      req.Name,
		Temp:      temp,
		ShelfLife: float64(req.ShelfLife),
		DecayRate: req.DecayRate,

		IsOnShelf:  make(chan struct{}, 1),
		IsCanceled: make(chan struct{}, 1),
	}, nil
}
