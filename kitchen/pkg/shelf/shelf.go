package shelf

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"sync"
	"time"
)

type shelf []*core.Order

func newShelf(capacity int) shelf {
	// create a shelf with max capacity: capacity
	return make([]*core.Order, 0, capacity)
}

func (shelf *shelf) add(order *core.Order) int {
	*shelf = append(*shelf, order)
	return len(*shelf) - 1
}

func (shelf *shelf) remove(idx int) *core.Order {
	ret := (*shelf)[idx]
	(*shelf)[idx] = (*shelf)[len(*shelf)-1]
	*shelf = (*shelf)[:len(*shelf)-1]
	return ret
}

func (shelf shelf) isFull() bool {
	return len(shelf) == cap(shelf)
}

func (shelf shelf) isEmpty() bool {
	return len(shelf) == 0
}

type shelves []shelf

func newShelves(capHot, capCold, capFrozen int) shelves {
	ret := make([]shelf, core.InvalidTemp)
	ret[core.Hot] = newShelf(capHot)
	ret[core.Cold] = newShelf(capCold)
	ret[core.Frozen] = newShelf(capFrozen)
	return ret
}

func (s shelves) remove(temp core.OrderTemp, idx int) *core.Order {
	return s[temp].remove(idx)
}

type singleTempShelves shelves

func newSingleTempShelves(capHot, capCold, capFrozen int) singleTempShelves {
	return singleTempShelves(newShelves(capHot, capCold, capFrozen))
}

type overFlowShelf shelves

func newOverFlowShelf(capacity int) overFlowShelf {
	return overFlowShelf(newShelves(capacity, capacity, capacity))
}

func (shelf overFlowShelf) add(order *core.Order) int {
	return shelf[order.Temp].add(order)
}

func (shelf overFlowShelf) isFull() bool {
	return len(shelf[core.Hot])+len(shelf[core.Cold])+len(shelf[core.Frozen]) == cap(shelf[core.Hot])
}

func (shelf overFlowShelf) remove(temp core.OrderTemp, idx int) *core.Order {
	return shelf[temp].remove(idx)
}

type shelfCat int

const (
	singleTemp shelfCat = iota
	overflow
)

type shelfLocation struct {
	cate shelfCat
	temp core.OrderTemp
	idx  int
}

type shelfMgr struct {
	ctx     *core.Context
	kitchen core.Kitchen

	singleShelves singleTempShelves
	overflowShelf overFlowShelf
	index         map[string]shelfLocation

	mutex sync.Mutex
}

func NewShelfMgr(ctx *core.Context) *shelfMgr {
	capHot := ctx.ShelfCap.Hot
	capCold := ctx.ShelfCap.Cold
	capFrozen := ctx.ShelfCap.Frozen
	capOverflow := ctx.ShelfCap.Overflow
	return &shelfMgr{
		ctx:           ctx,
		singleShelves: newSingleTempShelves(capHot, capCold, capFrozen),
		overflowShelf: newOverFlowShelf(capOverflow),
		index:         map[string]shelfLocation{},
	}
}

func (mgr *shelfMgr) SetKitchen(kitchen core.Kitchen) {
	mgr.kitchen = kitchen
}

func (mgr *shelfMgr) Put(order *core.Order) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	mgr.ctx.Log.Infof("put order %s, %p", order.ID, order)
	defer func() {
		mgr.ctx.Log.Infof("put order %s on %+v", order.ID, mgr.index[order.ID])
		//mgr.ctx.Log.Infof("index: %v", mgr.index)
	}()
	/*
		log := func(s string) {
			if order.ID == "68515b89-96bf-48b6-a07c-68cf17ca3c25" {
				mgr.ctx.Log.Info("%s", s)
			}
		}
	*/
	//id := "dcafe7ca-f7a7-4262-9b14-ab19729b2055"

	now := time.Now()
	temp := order.Temp
	order.UpdateTime = now
	order.RemainLife = order.ShelfLife

	// try placing on a shelf that matches the order’s temperature.
	if !mgr.singleShelves[temp].isFull() {
		idx := mgr.singleShelves[temp].add(order)
		mgr.index[order.ID] = shelfLocation{singleTemp, temp, idx}
		return
	}

	// try placing on the overflow shelf
	if !mgr.overflowShelf.isFull() {
		idx := mgr.overflowShelf.add(order)
		mgr.index[order.ID] = shelfLocation{overflow, temp, idx}
		return
	}

	// try moving an existing order on the overflow to an allowable shelf with room
	for temp = core.Hot; temp < core.InvalidTemp; temp++ {
		if !mgr.singleShelves[temp].isFull() &&
			!mgr.overflowShelf[temp].isEmpty() {

			toMove := mgr.overflowShelf.remove(temp, 0)
			if len(mgr.overflowShelf[temp]) > 0 {
				id := mgr.overflowShelf[temp][0].ID
				mgr.index[id] = shelfLocation{overflow, temp, 0}
			}
			toMove.UpdateRemainLife(now, true)
			idx := mgr.singleShelves[temp].add(toMove)
			mgr.index[toMove.ID] = shelfLocation{singleTemp, temp, idx}
			mgr.ctx.Log.Infof("move order %s", toMove.ID)
			idx = mgr.overflowShelf.add(order)
			mgr.index[order.ID] = shelfLocation{overflow, temp, idx}
			return
		}
	}

	// discard an order with least value from the overflow shelf
	minValue := 1.0
	var idx int
	var toDiscard *core.Order
	for temp = core.Hot; temp < core.InvalidTemp; temp++ {
		shelf := mgr.overflowShelf[temp]
		for i := 0; i < len(shelf); i++ {
			//mgr.ctx.Log.Infof("order: %+v", order)
			shelf[i].UpdateRemainLife(now, true)
			//mgr.ctx.Log.Infof("updated order: %+v", order)
			value := shelf[i].EstimatePickValue(true)
			//mgr.ctx.Log.Infof("value: %f", value)
			if value < minValue {
				minValue = value
				toDiscard = shelf[i]
				idx = i
			}
		}
	}
	//mgr.ctx.Log.Infof("toDiscard: %+v", toDiscard)
	mgr.ctx.Log.Infof("discard order: %s", toDiscard.ID)
	mgr.overflowShelf[toDiscard.Temp].remove(idx)
	delete(mgr.index, toDiscard.ID)
	if len(mgr.overflowShelf[toDiscard.Temp]) > idx {
		id := mgr.overflowShelf[toDiscard.Temp][idx].ID
		mgr.index[id] = shelfLocation{overflow, toDiscard.Temp, 0}
	}
	idx = mgr.overflowShelf.add(order)
	mgr.index[order.ID] = shelfLocation{overflow, order.Temp, idx}
	mgr.ctx.Log.Infof("orderID: %s", mgr.overflowShelf[order.Temp][idx].ID)
}

func (mgr *shelfMgr) Pick(orderID string) error {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	loc, found := mgr.index[orderID]
	if !found {
		// order has been discarded
		return core.ResourceNotFound.WithField("ID", orderID)
	}
	mgr.ctx.Log.Infof("Pick order %s from %+v", orderID, loc)
	if loc.cate == overflow {
		mgr.ctx.Log.Infof("overflow shelf: %v", mgr.overflowShelf)
		mgr.overflowShelf[loc.temp].remove(loc.idx)
		if len(mgr.overflowShelf[loc.temp]) > loc.idx {
			id := mgr.overflowShelf[loc.temp][loc.idx].ID
			mgr.ctx.Log.Infof("overflow move order %s from %v to %v", id, mgr.index[id], loc)
			mgr.index[id] = loc
		}
	} else {
		mgr.ctx.Log.Infof("singleshelf: %v", mgr.singleShelves)
		mgr.singleShelves[loc.temp].remove(loc.idx)
		if len(mgr.singleShelves[loc.temp]) > loc.idx {
			id := mgr.singleShelves[loc.temp][loc.idx].ID
			mgr.ctx.Log.Infof("singlTemp move order %s from %v to %v", id, mgr.index[id], loc)
			mgr.index[id] = loc
		}
	}
	delete(mgr.index, orderID)
	return nil
}
