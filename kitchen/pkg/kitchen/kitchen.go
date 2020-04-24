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
		ctx.Logger.WithError(err).Warn("invalid order request")
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
		k.orderChan = make(chan *core.Order, 1)
		k.stopChan = make(chan struct{}, 1)
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
					k.dispatch(order)
				case <-k.stopChan:
					break
				}
			}
		}
		close(k.stopChan)
		close(k.orderChan)
	}()
	return k
}

func (k *kitchen) Stop() {
	k.cookMgr.GetOffWork()
	k.courierMgr.GetOffWork()
	k.stopChan <- struct{}{}
}

func (k *kitchen) dispatch(order *core.Order) {
	switch order.Status {
	case core.Accepted:
		//
		k.cookMgr.Notify(order)
		k.courierMgr.Notify(order)
	case core.Cooked:
		// put onto shelve
	case core.Picking:
		// remove from shelve
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

		WaitCook: make(chan struct{}, 1),
		WaitPick: make(chan struct{}, 1),
	}, nil
}
