package kitchen

import (
	"kitchen/pkg/core"
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

type shelfSet struct {
	ctx     *core.Context
	kitchen core.Kitchen

	singleShelves singleTempShelves
	overflowShelf overflowShelf

	mutex sync.Mutex
}

func newShelfSet(ctx *core.Context) *shelfSet {
	isDebug = ctx.IsDebug
	capHot := ctx.ShelfCap.Hot
	capCold := ctx.ShelfCap.Cold
	capFrozen := ctx.ShelfCap.Frozen
	capOverflow := ctx.ShelfCap.Overflow
	return &shelfSet{
		ctx:           ctx,
		singleShelves: newSingleTempShelves(capHot, capCold, capFrozen),
		overflowShelf: newOverFlowShelf(capOverflow),
	}
}

func (s *shelfSet) setKitchen(kitchen core.Kitchen) {
	s.kitchen = kitchen
}

func (s *shelfSet) Put(order *core.Order) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ctx.Log.Infof("put order %s, %p", order.ID, order)

	now := time.Now()
	temp := order.Temp
	order.UpdateTime = now
	order.RemainLife = order.ShelfLife

	// try placing on a shelf that matches the order’s temperature.
	if !s.singleShelves[temp].isFull() {
		s.singleShelves[temp].add(order)
		order.ShelfType = core.SingleTempShelf
		return
	}

	// try placing on the overflow shelf
	if !s.overflowShelf.isFull() {
		s.overflowShelf.add(order)
		order.ShelfType = core.OverflowShelf
		return
	}

	// try moving an existing order on the overflow to an allowable shelf with room
	for t := core.Hot; t < core.InvalidTemp; t++ {
		if !s.overflowShelf[t].isEmpty() &&
			!s.singleShelves[t].isFull() {

			toMove := s.overflowShelf[t].pickFirst()
			updateRemainLife(toMove, now, true)
			s.overflowShelf.remove(toMove)
			s.singleShelves[toMove.Temp].add(toMove)
			toMove.ShelfType = core.SingleTempShelf
			s.kitchen.Send(toMove, core.Moved)
			s.overflowShelf.add(order)
			order.ShelfType = core.OverflowShelf
			return
		}
	}

	// discard an order with least value from the overflow shelf
	minValue := 1.0
	var toDiscard *core.Order
	for t := core.Hot; t < core.InvalidTemp; t++ {
		shelf := s.overflowShelf[t]
		for _, o := range shelf.content {
			updateRemainLife(o, now, true)
			if o.EstimatePickTime.Before(now) {
				o.EstimatePickTime = now.Add(time.Second)
			}
			value := o.EstimatePickValue(true)
			s.ctx.Log.Infof("estimate pick time: %v, %d", o.EstimatePickTime, o.EstimatePickTime.Unix())
			s.ctx.Log.Infof("estimate pick value: %f", value)
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
		s.ctx.Log.Error("all order value bigger than 1, impossible")
		panic("all order value bigger than 1")
	}
	s.ctx.Log.Infof("discard order: %s", toDiscard.ID)
	s.overflowShelf.remove(toDiscard)
	s.kitchen.Send(toDiscard, core.Discarded)
	s.overflowShelf.add(order)
	order.ShelfType = core.OverflowShelf
}

func (s *shelfSet) Pick(order *core.Order) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ctx.Log.Infof("pick order %s", order.ID)

	if s.singleShelves[order.Temp].find(order) {
		s.ctx.Log.Infof("order %+v picked", (s.singleShelves[order.Temp].content[order.ID]))
		s.singleShelves[order.Temp].remove(order)
	} else if s.overflowShelf.find(order) {
		s.ctx.Log.Infof("order %+v picked", (s.overflowShelf[order.Temp].content[order.ID]))
		s.overflowShelf.remove(order)
	} else {
		return core.ResourceNotFound.WithField("ID", order.ID)
	}
	return nil
}

func (s *shelfSet) content() [core.CountTempType + 1][]*core.Order {
	var ret [core.CountTempType + 1][]*core.Order
	now := time.Now()
	// single shelvs
	for t := core.Hot; t < core.CountTempType; t++ {
		ret[t] = make([]*core.Order, 0, s.singleShelves[t].size())
		for _, v := range s.singleShelves[t].content {
			updateRemainLife(v, now, false)
			ret[t] = append(ret[t], v)
		}
	}
	// overflow shelf
	overFlowIdx := core.CountTempType
	ret[overFlowIdx] = make([]*core.Order, 0, s.overflowShelf.size())
	for t := core.Hot; t < core.CountTempType; t++ {
		for _, v := range s.overflowShelf[t].content {
			updateRemainLife(v, now, true)
			ret[overFlowIdx] = append(ret[overFlowIdx], v)
		}
	}
	return ret
}

func updateRemainLife(o *core.Order, now time.Time, onOverFlow bool) {
	age := now.Unix() - o.UpdateTime.Unix()
	shelfDecayModifier := 1
	if onOverFlow {
		shelfDecayModifier = 2
	}
	o.UpdateTime = now
	o.RemainLife -= o.DecayRate * float64(age) * float64(shelfDecayModifier)
	if o.RemainLife < 0 {
		o.RemainLife = 0
	}
}
