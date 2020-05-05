package cook

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"time"
)

type cookMgr struct {
	*core.BaseColleague
	ctx *core.Context
}

func NewCookMgr(ctx *core.Context) *cookMgr {
	return &cookMgr{
		BaseColleague: core.NewBaseColleague(),
		ctx:           ctx,
	}
}

func (c *cookMgr) Notify(order *core.Order, event core.Event) {
	c.ctx.Log.Infof("cookMgr: receive Event %d for order %s", event, order.ID)
	go func() {
		c.cook(order)
		c.Kitchen.GetShelf().Put(order)
		c.Kitchen.Send(order, core.Cooked)
	}()
}

func (c *cookMgr) GetOffWork() {
	time.Sleep(time.Millisecond * 100)
}

func (c *cookMgr) cook(order *core.Order) {
	c.ctx.Log.Info("cook")
	time.Sleep(time.Millisecond * 50)
	order.Status = "cooked"
}
