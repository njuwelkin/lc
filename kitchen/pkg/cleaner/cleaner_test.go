package cleaner_test

import (
	. "github.com/onsi/gomega"
	//"sync/atomic"
	"testing"
	"time"

	"kitchen/pkg/cleaner"
	"kitchen/pkg/core"
	"kitchen/pkg/kitchen"
	"kitchen/pkg/test"
)

func TestCleaner(tt *testing.T) {
	t := test.NewTest(tt)

	req1 := &core.OrderRequest{ID: "h1", Temp: "hot", ShelfLife: 1, DecayRate: 0.6}
	//req2 := &core.OrderRequest{ID: "h2", Temp: "hot", ShelfLife: 1, DecayRate: 0.6}
	o1, _ := kitchen.NewOrder(req1)
	//o2, _ := kitchen.NewOrder(req2)

	k := test.NewKitchen(t.Context)
	c := cleaner.NewCleaner(t.Context)
	c.SetKitchen(k)

	t.Run("new clean job on single temp shelf", func() {
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Cooked)
		// o1 will be cleaned after (1/0.6) second
		time.Sleep(time.Second)
		found := k.Find(o1)
		t.Expect(found).To(BeTrue())
		time.Sleep(time.Second)
		found = k.Find(o1)
		t.Expect(found).To(BeFalse())

		// kitchen received a discard event
		t.Expect(k.Count).To(Equal(1))
	})

	t.Run("new clean job on overflow shelf", func() {
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.OverflowShelf
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Cooked)
		// o1 will be cleaned after (1/(0.6 * 2)) second
		time.Sleep(time.Millisecond * 500)
		found := k.Find(o1)
		t.Expect(found).To(BeTrue())
		time.Sleep(time.Millisecond * 500)
		found = k.Find(o1)
		t.Expect(found).To(BeFalse())

		// kitchen received a discard event
		t.Expect(k.Count).To(Equal(2))
	})

	t.Run("cancel a clean job on delieved", func() {
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Moved)
		// o1 will be cleaned after (1/ 0.6) second
		time.Sleep(time.Millisecond * 500)
		found := k.Find(o1)
		t.Expect(found).To(BeTrue())

		c.Notify(o1, core.Delivered)
		time.Sleep(time.Millisecond * 1500)
		found = k.Find(o1)
		t.Expect(found).To(BeTrue())

		// kitchen doesn't receive discard event
		t.Expect(k.Count).To(Equal(2))
	})

	t.Run("cancel a clean job on discarded", func() {
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Cooked)
		// o1 will be cleaned after (1/ 0.6) second
		time.Sleep(time.Millisecond * 500)
		found := k.Find(o1)
		t.Expect(found).To(BeTrue())

		c.Notify(o1, core.Discarded)
		time.Sleep(time.Millisecond * 1500)
		found = k.Find(o1)
		t.Expect(found).To(BeTrue())

		// kitchen doesn't receive discard event
		t.Expect(k.Count).To(Equal(2))
	})

	t.Run("all clean job should be done before cleaner getoffwork", func() {
		o1.UpdateTime = time.Now()
		o1.ShelfType = core.SingleTempShelf
		k.GetShelf().Put(o1)
		c.Notify(o1, core.Cooked)
		// o1 will be cleaned after (1/0.6) second
		c.GetOffWork()
		found := k.Find(o1)
		t.Expect(found).To(BeFalse())

		// kitchen received a discard event
		t.Expect(k.Count).To(Equal(3))
	})

}
