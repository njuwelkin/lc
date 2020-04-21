package shelf

import (
	"github.com/njuwelkin/lc/orders/pkg/core"
	"sync"
)

type shelf struct {
	content []*core.Order
	used    int
}

func newShelf(capacity int) {
	return &shelf{
		content:  make([]*core.Order, capacity),
		occupied: 0,
	}
}

func (shelf *shelf) add(order *core.Order) {
	shelf.content[shelf.used] = order
	used++
}

func (shelf *shelf) remove(idx int) *core.Order {
	ret := shelf.content[idx]
	shelf.content[idx] = shelf.content[shelf.used-1]
	shelf.used--
	return ret
}

func (shelf *shelf) isFull() bool {
	return shelf.used == len(shelf.content)
}

func (shelf *shelf) isEmpty() bool {
	return shelf.used == 0
}

type shelves []*shelf

func newShelves(capHot, capCold, capFrozen int) *singleTempShelves {
	ret := singleTempShelves(make([]*shelf, core.InvalidTemp))
	ret[core.hot] = newshelf(caphot)
	ret[core.cold] = newshelf(capcold)
	ret[core.frozen] = newshelf(capfrozen)
	return &ret
}

func (s *shelves) remove(temp core.OrderTemp, idx int) *core.Order {
	return s[temp].remove(idx)
}

type sigleTempShelves shelves

func newSingleTempShelves(capHot, capCold, capFrozen int) *sigleTempShelves {
	return newShelves(capHot, capCold, capFrozen)
}

type overFlowShelf shelves

func newOverFlowShelf(capacity int) *overFlowShelf {
	return newShelves(capacity, capacity, capacity)
}

func (shelf *overFlowShelf) add(order *core.Order) bool {
	if shelf.isFull() {
		return false
	}
	return ret.content[order.GetTemp()].add(order)
}

func (shelf *overFlowShelf) isFull() bool {
	return shelf.hot.used+shelf.cold.used+shelf.frozen.used == len(hot.content)
}

type shelfMgr struct {
	ctx           *core.Context
	singleShelves *singleTempShelves
	overflowShelf *overFlowShelf

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
	}
}

func (mgr *shelfMgr) Put(order *core.Order) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	temp := order.GetTemp()
	if !mgr.singleShelves[temp].isFull() {
		mgr.singleShelves[temp].add(order)
		return
	}
	if !mgr.overflowShelf.isFull() {
		mgr.overflowShelf.add(order)
		return
	}
	for temp = core.Hot; temp < core.InvalidTemp; temp++ {
		if !mgr.singleShelves[temp].isFull() &&
			!mgr.overflowShelf[temp].isEmpty() {

			tmp := mgr.overflowShelf.remove(temp, 0)
			mgr.singleShelves[temp].add(tmp)
			mgr.overflowShelf.add(order)
			return
		}
	}
	// discard one
	for temp = core.Hot; temp < core.InvalidTemp; temp++ {
		shelf := mgr.overflowShelf[temp]
	}
}

func (mgr *shelfMgr) Pick(order *core.Order) error {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	return nil
}
