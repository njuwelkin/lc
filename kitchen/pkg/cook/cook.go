package cook

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
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

func (c *cookMgr) Notify(order *core.Order) {
	c.cook(order)
	c.Kitchen.GetShelf().Put(order)
	c.Kitchen.Send(order, core.Cooked)
}

func (c *cookMgr) GetOffWork() {
}

func (c *cookMgr) cook(order *core.Order) {
	c.ctx.Log.Info("cook")
	order.Status = "cooked"
}
