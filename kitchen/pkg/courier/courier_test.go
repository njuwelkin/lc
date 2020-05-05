package courier_test

import (
	. "github.com/onsi/gomega"
	//"sync/atomic"
	"testing"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"github.com/njuwelkin/lc/kitchen/pkg/courier"
	"github.com/njuwelkin/lc/kitchen/pkg/kitchen"
	"github.com/njuwelkin/lc/kitchen/pkg/test"
)

func TestCourier(tt *testing.T) {
	t := test.NewTest(tt)
	// courier runs faster in ut
	t.MinPickDuration = 1
	t.MaxPickDuration = 2
	// two couriers
	t.NumOfCouriers = 2

	req1 := &core.OrderRequest{ID: "h1", Temp: "hot", ShelfLife: 2, DecayRate: 0.5}
	o1, _ := kitchen.NewOrder(req1)

	k := test.NewKitchen(t.Context)
	c := courier.NewCourierMgr(t.Context)
	c.SetKitchen(k)

	t.Run("pick an order", func() {
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		o1.Ready <- struct{}{}
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Accepted)
		// o1 will be picked after 1-2 second
		time.Sleep(time.Millisecond * 900)
		found := k.Find(o1)
		t.Expect(found).To(BeTrue())
		time.Sleep(time.Millisecond * 1200)
		found = k.Find(o1)
		t.Expect(found).To(BeFalse())

		// kitchen received a picked event and a delieverd event
		t.Expect(k.Count).To(Equal(2))
	})

	t.Run("abort task when the order is canceled", func() {
		k.Count = 0
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		o1.Ready <- struct{}{}
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Accepted)
		// o1 will be picked after 1-2 second
		time.Sleep(time.Millisecond * 900)
		found := k.Find(o1)
		t.Expect(found).To(BeTrue())
		// order is discarded
		o1.Cancel <- struct{}{}
		time.Sleep(time.Millisecond * 1200)
		// not picked
		found = k.Find(o1)
		t.Expect(found).To(BeTrue())
		// kitchen will not receive any event
		t.Expect(k.Count).To(Equal(0))
		// clean
		k.GetShelf().Pick(o1)
		_, _ = <-o1.Ready
	})

	t.Run("courier blocks if food is not ready", func() {
		k.Count = 0
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf

		c.Notify(o1, core.Accepted)
		// o1 will be picked after 1-2 second
		// food is not ready
		found := k.Find(o1)
		t.Expect(found).To(BeFalse())
		// courier arrives at kitchen and wait ther
		time.Sleep(time.Second * 2)
		t.Expect(k.Count).To(Equal(0))

		// food is ready
		k.GetShelf().Put(o1)
		o1.Ready <- struct{}{}
		// pick
		time.Sleep(time.Millisecond * 100)
		found = k.Find(o1)
		t.Expect(found).To(BeFalse())
		// kitchen recieved picked and delievered
		t.Expect(k.Count).To(Equal(2))
	})

	t.Run("all courier job should be done before get off work", func() {
		k.Count = 0
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		k.GetShelf().Put(o1)
		o1.Ready <- struct{}{}
		c.Notify(o1, core.Accepted)
		c.GetOffWork()
		t.Expect(k.Find(o1)).To(BeFalse())
		t.Expect(k.Count).To(Equal(2))
	})
}
