package kitchen

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
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

	shelf *shelfSet

	countDelieve int
	countDiscard int
}

func NewKitchen(ctx *core.Context, cookMgr, courierMgr core.Colleague) *kitchen {
	k := &kitchen{
		ctx:        ctx,
		msgQue:     make(chan *message, 4096),
		stop:       make(chan struct{}),
		cookMgr:    cookMgr,
		courierMgr: courierMgr,
		shelf:      newShelfSet(ctx),
	}
	cookMgr.SetKitchen(k)
	courierMgr.SetKitchen(k)
	k.shelf.setKitchen(k)

	return k
}

func (k *kitchen) PlaceOrder(req *core.OrderRequest) error {
	k.ctx.Log.Infof("receive order %+v", req)
	order, err := newOrder(req)
	if err != nil {
		k.ctx.Log.WithError(err).Warn("invalid order request")
		return err
	}
	k.Send(order, core.Accept)
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

func (k *kitchen) Stop() {
	// cooks complete existing job
	k.cookMgr.GetOffWork()
	// couriers deliver all existing orders
	k.courierMgr.GetOffWork()
	// leave one second for the out put
	time.Sleep(time.Second)
	k.stop <- struct{}{}
	k.ctx.Log.Infof("kitchen stopped, %d orders delivered, %d discarded", k.countDelieve, k.countDiscard)
	k.ctx.PrintStatistic(k.countDelieve, k.countDiscard)
}

func (k *kitchen) GetShelf() core.Shelf {
	return k.shelf
}

func (k *kitchen) dispatch(order *core.Order, event core.Event) {
	switch event {
	case core.Accept:
		// notify cook manager to prepare the food
		k.cookMgr.Notify(order, event)
		// notify courier manager to dispatch a courier
		k.courierMgr.Notify(order, event)
	case core.Cooked:
		// food is ready, wakeup waiting courier
		//order.IsOnShelf <- struct{}{}
		// notify cleaner to schedule a clean job
	case core.Moved:
		// notify cleaner to re-schedule the clean job
	case core.Discarded:
		k.courierMgr.Notify(order, event)
		k.countDiscard++
	case core.Delivered:
		k.countDelieve++
		k.ctx.Log.Infof("countDelivered %s: %d->%d", order.ID, k.countDelieve-1, k.countDelieve)
	}
	k.ctx.PrintEvent(order, event)
	//k.ctx.PrintShelfContent(k.shelf.content())
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
		ShelfLife:  float64(req.ShelfLife),
		RemainLife: float64(req.ShelfLife),
		DecayRate:  req.DecayRate,
		UpdateTime: time.Now(),
	}, nil
}
