package kitchen

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

type kitchen struct {
	ctx   *core.Context
	clerk core.Colleague
}

func NewKitchen(ctx *core.Context,
	clerk, cookMgr, courierMgr core.Colleague) *kitchen {

	clerk.AddSuccessor(cookMgr)
	clerk.AddSuccessor(courierMgr)
	return &kitchen{
		ctx:   ctx,
		clerk: clerk,
	}
}

func (k *kitchen) PlaceOrder(req *core.OrderRequest) error {
	k.ctx.Log.WithField("ID", req.ID).Info("receive order")
	order, err := newOrder(req)
	if err != nil {
		return err
	}
	k.clerk.Notify(order)
	return nil
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
	}, nil
}
