package shelf2

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"sync"
	"time"
)

var isDebug bool = false

type shelf struct {
	content  map[string]*core.Order
	capacity int
}

func newShelf(capacity int) *shelf {
	// create a shelf with max capacity: capacity
	return &shelf{
		content:  make(map[string]*core.Order),
		capacity: capacity,
	}
}

func (s *shelf) add(order *core.Order) {
	if isDebug && len(s.content) == s.capacity {
		// caller responsible for checking capacity
		panic("")
	}
	s.content[order.ID] = order
	return
}

func (s *shelf) find(order *core.Order) bool {
	return s.content[order.ID] != nil
}

func (s *shelf) remove(order *core.Order) {
	_, found := s.content[order.ID]
	if isDebug && !found {
		// caller responsible for checking existance
		panic("")
	}
	delete(s.content, order.ID)
}

func (s shelf) size() int {
	return len(s.content)
}

func (s shelf) cap() int {
	return s.capacity
}

func (s shelf) isFull() bool {
	return s.size() == s.cap()
}

func (s shelf) isEmpty() bool {
	return s.size() == 0
}

func (s shelf) pickFirst() *core.Order {
	for _, order := range s.content {
		return order
	}
	return nil
}

type singleTempShelves []*shelf

func newSingleTempShelves(capHot, capCold, capFrozen int) singleTempShelves {
	ret := make([]*shelf, core.InvalidTemp)
	ret[core.Hot] = newShelf(capHot)
	ret[core.Cold] = newShelf(capCold)
	ret[core.Frozen] = newShelf(capFrozen)
	return ret
}

type overflowShelf []*shelf

func newOverFlowShelf(capacity int) overflowShelf {
	ret := make([]*shelf, core.InvalidTemp)
	ret[core.Hot] = newShelf(capacity)
	ret[core.Cold] = newShelf(capacity)
	ret[core.Frozen] = newShelf(capacity)
	return ret
}

func (os overflowShelf) size() int {
	return os[core.Hot].size() + os[core.Cold].size() + os[core.Frozen].size()
}

func (os overflowShelf) cap() int {
	return os[core.Hot].cap()
}

func (os overflowShelf) add(order *core.Order) {
	if isDebug && os.size() == os.cap() {
		// caller responsible for checking capacity
		panic("")
	}
	os[order.Temp].content[order.ID] = order
	return
}

func (os overflowShelf) find(order *core.Order) bool {
	return os[order.Temp].find(order)
}

func (os overflowShelf) remove(order *core.Order) {
	os[order.Temp].remove(order)
}

func (os overflowShelf) isFull() bool {
	return os.size() == os.cap()
}

type shelfMgr struct {
	ctx     *core.Context
	kitchen core.Kitchen

	singleShelves singleTempShelves
	overflowShelf overflowShelf

	mutex sync.Mutex
}

func NewShelfMgr(ctx *core.Context) *shelfMgr {
	isDebug = ctx.IsDebug
	capHot := ctx.ShelfCap.Hot
	capCold := ctx.ShelfCap.Cold
	capFrozen := ctx.ShelfCap.Frozen
	capOverflow := ctx.ShelfCap.Overflow
	return &shelfMgr{
		ctx:           ctx,
		singleShelves: newSingleTempShelves(capHot, capCold, capFrozen),
		overflowShelf: newOverFlowShelf(capOverflow),
	}
}

func (mgr *shelfMgr) SetKitchen(kitchen core.Kitchen) {
	mgr.kitchen = kitchen
}

func (mgr *shelfMgr) Put(order *core.Order) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	mgr.ctx.Log.Infof("put order %s, %p", order.ID, order)

	now := time.Now()
	temp := order.Temp
	order.UpdateTime = now
	order.RemainLife = order.ShelfLife

	// try placing on a shelf that matches the order’s temperature.
	if !mgr.singleShelves[temp].isFull() {
		mgr.singleShelves[temp].add(order)
	}

	// try placing on the overflow shelf
	if !mgr.overflowShelf.isFull() {
		mgr.overflowShelf.add(order)
		return
	}

	// try moving an existing order on the overflow to an allowable shelf with room
	for t := core.Hot; t < core.InvalidTemp; t++ {
		if !mgr.overflowShelf[t].isEmpty() &&
			!mgr.singleShelves[t].isFull() {

			toMove := mgr.overflowShelf[t].pickFirst()
			toMove.UpdateRemainLife(now, true)
			mgr.overflowShelf.remove(toMove)
			mgr.singleShelves[toMove.Temp].add(toMove)
			mgr.overflowShelf.add(order)
			return
		}
	}

	// discard an order with least value from the overflow shelf
	minValue := 1.0
	var toDiscard *core.Order
	for t := core.Hot; t < core.InvalidTemp; t++ {
		shelf := mgr.overflowShelf[t]
		for _, o := range shelf.content {
			o.UpdateRemainLife(now, true)
			if o.EstimatePickTime.Before(now) {
				o.EstimatePickTime = now.Add(time.Second)
			}
			value := o.EstimatePickValue(true)
			mgr.ctx.Log.Infof("estimate pick time: %v, %d", o.EstimatePickTime, o.EstimatePickTime.Unix())
			mgr.ctx.Log.Infof("estimate pick value: %f", value)
			if value <= minValue {
				minValue = value
				toDiscard = o
				if value == 0 {
					break
				}
			}
		}
	}
	if toDiscard == nil {
		mgr.ctx.Log.Error("all order value bigger than 1, impossible")
		panic("all order value bigger than 1")
	}
	mgr.ctx.Log.Infof("discard order: %s", toDiscard.ID)
	mgr.overflowShelf.remove(toDiscard)
	mgr.kitchen.Send(toDiscard, core.Discarded)
	mgr.overflowShelf.add(order)
}

func (mgr *shelfMgr) Pick(order *core.Order) error {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	mgr.ctx.Log.Infof("pick order %s", order.ID)

	if mgr.singleShelves[order.Temp].find(order) {
		mgr.ctx.Log.Infof("order %+v picked", (mgr.singleShelves[order.Temp].content[order.ID]))
		mgr.singleShelves[order.Temp].remove(order)
	} else if mgr.overflowShelf.find(order) {
		mgr.ctx.Log.Infof("order %+v picked", (mgr.overflowShelf[order.Temp].content[order.ID]))
		mgr.overflowShelf.remove(order)
	} else {
		return core.ResourceNotFound.WithField("ID", order.ID)
	}
	return nil
}
