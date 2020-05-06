package kitchen

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"

	"kitchen/pkg/core"
	"kitchen/pkg/test"
)

type mockKitchen struct {
	countMoveEvent    int
	countDiscardEvent int
}

func (k *mockKitchen) Send(o *core.Order, e core.Event) {
	switch e {
	case core.Moved:
		k.countMoveEvent++
	case core.Discarded:
		k.countDiscardEvent++
	}
}

func (k *mockKitchen) GetShelf() core.Shelf {
	return nil
}

func TestShelfSet(tt *testing.T) {
	t := test.NewTest(tt)
	t.ShelfCap.Hot = 1
	t.ShelfCap.Cold = 1
	t.ShelfCap.Frozen = 1
	t.ShelfCap.Overflow = 2

	h1, _ := NewOrder(&core.OrderRequest{ID: "h1", Temp: "hot", ShelfLife: 20, DecayRate: 0.6})
	h2, _ := NewOrder(&core.OrderRequest{ID: "h2", Temp: "hot", ShelfLife: 20, DecayRate: 0.3})
	c1, _ := NewOrder(&core.OrderRequest{ID: "c1", Temp: "cold", ShelfLife: 20, DecayRate: 0.6})
	c2, _ := NewOrder(&core.OrderRequest{ID: "c2", Temp: "cold", ShelfLife: 20, DecayRate: 0.3})
	f1, _ := NewOrder(&core.OrderRequest{ID: "f1", Temp: "frozen", ShelfLife: 20, DecayRate: 0.6})
	f2, _ := NewOrder(&core.OrderRequest{ID: "f2", Temp: "frozen", ShelfLife: 20, DecayRate: 0.3})
	h3, _ := NewOrder(&core.OrderRequest{ID: "h3", Temp: "hot", ShelfLife: 20, DecayRate: 0.3})

	s := newShelfSet(t.Context)
	k := mockKitchen{}
	s.setKitchen(&k)

	// Each order should be placed on a shelf that matches the order’s temperature.
	t.Run("put food on single temp shelf", func() {
		s.Put(h1)
		t.Expect(s.singleShelves[core.Hot].size()).To(Equal(1))
		t.Expect(h1.ShelfType).To(Equal(core.SingleTempShelf))
		s.Put(c1)
		t.Expect(s.singleShelves[core.Cold].size()).To(Equal(1))
		s.Put(f1)
		t.Expect(s.singleShelves[core.Frozen].size()).To(Equal(1))
	})

	// If that shelf is full, an order can be placed on the overflow shelf.
	t.Run("put food on overflow shelf", func() {
		s.Put(h2)
		t.Expect(s.overflowShelf.size()).To(Equal(1))
		t.Expect(s.overflowShelf[core.Hot].size()).To(Equal(1))
		t.Expect(h2.ShelfType).To(Equal(core.OverflowShelf))
		s.Put(c2)
		t.Expect(s.overflowShelf.size()).To(Equal(2))
		t.Expect(s.overflowShelf[core.Cold].size()).To(Equal(1))
	})

	// If the overflow shelf is full, an existing order of your choosing on the overflow
	//    should be moved to an allowable shelf with room
	t.Run("move food from overflow shelf to single temp shelf", func() {
		err := s.Pick(h1)
		t.Expect(err).NotTo(HaveOccurred())
		t.Expect(s.singleShelves[core.Hot].size()).To(Equal(0))

		// both frozen shelf and overflow shelf is full
		//    but hot shelf is not full
		s.Put(f2)
		// h2 is moved to single temp shelf
		t.Expect(s.singleShelves[core.Hot].size()).To(Equal(1))
		t.Expect(s.singleShelves[core.Hot].find(h2)).To(BeTrue())
		t.Expect(h2.ShelfType).To(Equal(core.SingleTempShelf))
		t.Expect(s.overflowShelf.size()).To(Equal(2))
		t.Expect(s.overflowShelf[core.Cold].size()).To(Equal(1))
		t.Expect(f2.ShelfType).To(Equal(core.OverflowShelf))
		// kitchen has received the move event
		t.Expect(k.countMoveEvent).To(Equal(1))
	})

	// If no such move is possible, an order from the overflow shelf should be discarded as waste
	t.Run("discard food on overflow shelf", func() {
		// now all sheves are full
		t.Expect(s.singleShelves[core.Hot].isFull()).To(BeTrue())
		t.Expect(s.singleShelves[core.Cold].isFull()).To(BeTrue())
		t.Expect(s.singleShelves[core.Frozen].isFull()).To(BeTrue())
		t.Expect(s.overflowShelf.isFull()).To(BeTrue())

		// now order f2 and c2 are on overflow shelf
		t.Expect(s.overflowShelf.find(f2)).To(BeTrue())
		t.Expect(s.overflowShelf.find(c2)).To(BeTrue())

		// f2's estimate pick value is more than c2
		now := time.Now()
		f2.EstimatePickTime = now.Add(time.Second * 2)
		c2.EstimatePickTime = now.Add(time.Second * 3)

		// now put a new order, c2 will be discarded
		//    and h3 will be put on overflow shelf
		s.Put(h3)
		t.Expect(h3.ShelfType).To(Equal(core.OverflowShelf))
		t.Expect(s.overflowShelf.find(c2)).To(BeFalse())
		// kitchen received the discard event
		t.Expect(k.countDiscardEvent).To(Equal(1))

		err := s.Pick(c2)
		t.Expect(err).To(HaveOccurred())
		t.Expect(core.ResourceNotFound.Is(err)).To(BeTrue())
	})

}
