package kitchen

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

type kitchen struct {
	// context
	ctx *core.Context

	// members
	cookMgr    core.Colleague
	courierMgr core.Colleague

	orderChan chan *core.Order
	stopChan  chan struct{}
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
	k.orderChan <- order
}

func (k *kitchen) Run() *kitchen {
	go func() {
		k.orderChan = make(chan *core.Order, 10000)
		k.stopChan = make(chan struct{})
		defer func() {
			close(k.stopChan)
			close(k.orderChan)
		}()
		for {
			// two layer select to make sure order's
			//   priority is higher than stop.
			// in order that
			select {
			case order := <-k.orderChan:
				k.dispatch(order)
			default:
				select {
				case order := <-k.orderChan:
					k.ctx.Log.Info("receive order msg")
					k.dispatch(order)
				case <-k.stopChan:
					k.ctx.Log.Info("receive stop msg")
					return
				}
			}
		}
	}()
	return k
}

func (k *kitchen) Stop() {
	k.cookMgr.GetOffWork()
	k.courierMgr.GetOffWork()
	k.stopChan <- struct{}{}
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
		// notify shelf cleaner to schedule a clean job
	case core.Picked:
		// notify shelf cleaner to remove the clean job
	case core.Discarded:
	case core.Delivered:
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

		//WaitCook: make(chan struct{}, 1),
	}, nil
}
