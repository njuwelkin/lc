package clerk

import (
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

type clerk struct {
	*core.BaseColleague
	ctx *core.Context
}

func NewClerk(ctx *core.Context) *clerk {
	return &clerk{
		BaseColleague: core.NewBaseColleague(),
		ctx:           ctx,
	}
}

func (c *clerk) Notify(order *core.Order) {
	c.Send(order)
}
