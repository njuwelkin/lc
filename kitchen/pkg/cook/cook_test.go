package cook_test

import (
	. "github.com/onsi/gomega"
	//"sync/atomic"
	"testing"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/cook"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"github.com/njuwelkin/lc/kitchen/pkg/kitchen"
	"github.com/njuwelkin/lc/kitchen/pkg/test"
)

func TestCook(tt *testing.T) {
	t := test.NewTest(tt)

	req1 := &core.OrderRequest{ID: "h1", Temp: "hot", ShelfLife: 2, DecayRate: 0.5}
	o1, _ := kitchen.NewOrder(req1)

	k := test.NewKitchen(t.Context)
	c := cook.NewCookMgr(t.Context)
	c.SetKitchen(k)

	t.Run("pick an order", func() {
		c.Notify(o1, core.Accepted)
		time.Sleep(time.Millisecond * 100)
		t.Expect(o1.Status).To(Equal("cooked"))
		t.Expect(k.Find(o1)).To(BeTrue())
		// kitchen received cooked event
		t.Expect(k.Count).To(Equal(1))
	})

}
