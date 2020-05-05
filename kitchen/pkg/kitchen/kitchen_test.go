package kitchen_test

import (
	. "github.com/onsi/gomega"
	//"sync/atomic"
	"testing"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"github.com/njuwelkin/lc/kitchen/pkg/kitchen"
	"github.com/njuwelkin/lc/kitchen/pkg/test"
)

type colleague struct {
	count   [core.CountEvent]int
	stopped bool
}

func newColleague() *colleague {
	return &colleague{}
}

func (c *colleague) SetKitchen(k core.Kitchen) {
}

func (c *colleague) Notify(o *core.Order, e core.Event) {
	c.count[e]++
}

func (c *colleague) GetOffWork() {
	c.stopped = true
}

func TestKitchen(tt *testing.T) {
	t := test.NewTest(tt)

	req1 := &core.OrderRequest{ID: "h1", Temp: "invalid temp", ShelfLife: 20, DecayRate: 0.6}
	req2 := &core.OrderRequest{ID: "h2", Temp: "hot", ShelfLife: 20, DecayRate: 0.6}
	order1, _ := kitchen.NewOrder(req2)
	order2, _ := kitchen.NewOrder(req2)

	cookMgr := newColleague()
	courierMgr := newColleague()
	cleaner := newColleague()
	k := kitchen.NewKitchen(t.Context, cookMgr, courierMgr, cleaner)
	k.Run()

	t.Run("place a new order", func() {
		err := k.PlaceOrder(req1)
		t.Expect(err).To(HaveOccurred())
		t.Expect(core.InvalidOrderRequest.Is(err)).To(BeTrue())
	})
	t.Run("place a new order", func() {
		err := k.PlaceOrder(req2)
		t.Expect(err).To(BeNil())
		time.Sleep(time.Millisecond * 10)
		// cookMgr received the event
		t.Expect(cookMgr.count[core.Accepted]).To(Equal(1))
		// courierMgr received the event
		t.Expect(courierMgr.count[core.Accepted]).To(Equal(1))
	})

	t.Run("send cooked event", func() {
		k.Send(order1, core.Cooked)
		time.Sleep(time.Millisecond * 10)
		// cleaner received the event
		t.Expect(cleaner.count[core.Cooked]).To(Equal(1))
		// order is ready for picking
		_, ok := <-order1.Ready
		t.Expect(ok).To(BeTrue())
	})

	t.Run("send moved event", func() {
		k.Send(order1, core.Moved)
		time.Sleep(time.Millisecond * 10)
		// cleaner received the event
		t.Expect(cleaner.count[core.Moved]).To(Equal(1))
	})

	t.Run("send discard event", func() {
		k.Send(order1, core.Discarded)
		time.Sleep(time.Millisecond * 10)
		// cleaner received the event
		t.Expect(cleaner.count[core.Discarded]).To(Equal(1))
	})

	t.Run("send delieved event", func() {
		k.Send(order2, core.Delivered)
		time.Sleep(time.Millisecond * 10)
		// cleaner received the event
		t.Expect(cleaner.count[core.Delivered]).To(Equal(1))
	})

	t.Run("stop kitchen", func() {
		k.Stop()
		t.Expect(cookMgr.stopped).To(BeTrue())
		t.Expect(courierMgr.stopped).To(BeTrue())
		t.Expect(cleaner.stopped).To(BeTrue())
	})
}
