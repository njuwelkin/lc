package kitchen

import (
	"kitchen/pkg/core"
	"time"
)

type message struct {
	order *core.Order
	event core.Event
}

type kitchen struct {
	// context
	ctx    *core.Context
	msgQue chan *message
	stop   chan struct{}

	// colleagues
	cookMgr    core.Colleague
	courierMgr core.Colleague
	cleaner    core.Colleague

	shelf *shelfSet

	countDelieve int
	countDiscard int
}

func NewKitchen(ctx *core.Context, cookMgr, courierMgr, cleaner core.Colleague) *kitchen {
	k := &kitchen{
		ctx:        ctx,
		msgQue:     make(chan *message, 4096),
		stop:       make(chan struct{}),
		cookMgr:    cookMgr,
		courierMgr: courierMgr,
		cleaner:    cleaner,
		shelf:      newShelfSet(ctx),
	}
	cookMgr.SetKitchen(k)
	courierMgr.SetKitchen(k)
	cleaner.SetKitchen(k)
	k.shelf.setKitchen(k)

	return k
}

func (k *kitchen) PlaceOrder(req *core.OrderRequest) error {
	k.ctx.Log.Infof("receive order %+v", req)
	order, err := NewOrder(req)
	if err != nil {
		k.ctx.Log.WithError(err).Warn("invalid order request")
		return err
	}
	k.Send(order, core.Accepted)
	return nil
}

func (k *kitchen) Send(order *core.Order, event core.Event) {
	k.msgQue <- &message{order, event}
}

func (k *kitchen) run() {
	for {
		select {
		case msg := <-k.msgQue:
			k.dispatch(msg.order, msg.event)
		default:
			select {
			case msg := <-k.msgQue:
				k.dispatch(msg.order, msg.event)
			case <-k.stop:
				return
			}
		}
	}
}

func (k *kitchen) Run() *kitchen {
	go k.run()
	return k
}

// caller will block until all jobs done
func (k *kitchen) Stop() {
	// cooks complete existing job
	k.cookMgr.GetOffWork()
	// couriers deliver all existing orders
	k.courierMgr.GetOffWork()
	// clean all discarded food on shelf
	k.cleaner.GetOffWork()
	// stop the message loop
	k.stop <- struct{}{}
	k.ctx.PrintStatistic(k.countDelieve, k.countDiscard)
}

func (k *kitchen) GetShelf() core.Shelf {
	return k.shelf
}

func (k *kitchen) dispatch(order *core.Order, event core.Event) {
	switch event {
	case core.Accepted:
		// notify cook manager to prepare the food
		k.cookMgr.Notify(order, event)
		// notify courier manager to dispatch a courier
		k.courierMgr.Notify(order, event)
	case core.Cooked:
		// tell courier food is ready
		order.Ready <- struct{}{}
		// notify cleaner to schedule a clean job
		k.cleaner.Notify(order, event)
	case core.Moved:
		// notify cleaner to re-schedule the clean job
		k.cleaner.Notify(order, event)
	case core.Picked:
		// do nothing.
		// this event is not really need, as once it's picked,
		//   it will be delivered at once.
	case core.Discarded:
		// notify cleaner to remove the clean job
		k.cleaner.Notify(order, event)
		// tell courier to cancel the picking job
		order.Cancel <- struct{}{}
		// update statistic and close the order
		k.countDiscard++
		closeOrder(order)
	case core.Delivered:
		// notify cleaner to remove the clean job
		k.cleaner.Notify(order, event)
		// update statistic and close the order
		k.countDelieve++
		closeOrder(order)
		k.ctx.Log.Infof("countDelivered %s: %d->%d", order.ID, k.countDelieve-1, k.countDelieve)
	default:
		k.ctx.Log.Errorf("kitchen received an invalid event %d for order %s", event, order.ID)
		if k.ctx.IsDebug {
			panic("kitchen received an invalid event")
		}
	}
	k.ctx.PrintEvent(order, event)
	k.ctx.PrintShelfContent(k.shelf.content())
}

var tempNames = map[string]core.OrderTemp{
	"hot":    core.Hot,
	"cold":   core.Cold,
	"frozen": core.Frozen,
}

func NewOrder(req *core.OrderRequest) (*core.Order, error) {
	temp, found := tempNames[req.Temp]
	if !found {
		return nil, core.InvalidOrderRequest.WithField("Temp", req.Temp)
	}
	return &core.Order{
		ID:         req.ID,
		Name:       req.Name,
		Temp:       temp,
		ShelfLife:  float64(req.ShelfLife),
		RemainLife: float64(req.ShelfLife),
		DecayRate:  req.DecayRate,
		UpdateTime: time.Now(),

		Ready:  make(chan struct{}, 1),
		Cancel: make(chan struct{}, 1),
	}, nil
}

func closeOrder(order *core.Order) {
	close(order.Ready)
	close(order.Cancel)
}
