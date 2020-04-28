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

	countDelieve int
	countDiscard int
}

func NewKitchen(ctx *core.Context, cookMgr, courierMgr core.Colleague) *kitchen {
	k := &kitchen{
		ctx:        ctx,
		cookMgr:    cookMgr,
		courierMgr: courierMgr,
	}
	cookMgr.SetKitchen(k)
	courierMgr.SetKitchen(k)
	return k
}

func (k *kitchen) PlaceOrder(req *core.OrderRequest) error {
	k.ctx.Log.WithField("ID", req.ID).Info("receive order")
	order, err := newOrder(req)
	if err != nil {
		k.ctx.Log.WithError(err).Warn("invalid order request")
		return err
	}
	k.Send(order)
	return nil
}

func (k *kitchen) Send(order *core.Order) {
	k.dispatch(order)
}

func (k *kitchen) Run() *kitchen {
	return k
}

func (k *kitchen) Stop() {
	k.cookMgr.GetOffWork()
	k.courierMgr.GetOffWork()
}

func (k *kitchen) GetShelf() core.Shelf {
	return nil
}

func (k *kitchen) dispatch(order *core.Order) {
	switch order.Status {
	case core.Accepted:
		// notify cook manager to process the order
		k.cookMgr.Notify(order)
		// notify courier manager to dispatch a courier
		k.courierMgr.Notify(order)
	case core.Cooked:
		// wakeup waiting courier
		order.IsOnShelf <- struct{}{}
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
		ID:         req.ID,
		Name:       req.Name,
		Temp:       temp,
		ShelfLife:  req.ShelfLife,
		RemainLefe: req.ShelfLife,

		IsOnShelf: make(chan struct{}, 1),
	}, nil
}
