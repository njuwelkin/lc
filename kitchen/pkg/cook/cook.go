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
	c.Kitchen.Send(order)
}

func (c *cookMgr) cook(order *core.Order) {
	order.Status = core.Cooked
}
