package cleaner

import (
	"github.com/njuwelkin/lc/kitchen/pkg/common/scheduler"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"time"
)

type cleaner struct {
	*core.BaseColleague
	ctx       *core.Context
	scheduler *scheduler.Scheduler
}

func NewCleaner(ctx *core.Context) *cleaner {
	return &cleaner{
		BaseColleague: core.NewBaseColleague(),
		ctx:           ctx,
		scheduler:     scheduler.NewScheduler().Run(),
	}
}

func (c *cleaner) Notify(order *core.Order, event core.Event) {
	c.ctx.Log.Infof("cleaner: receive event %d for order %s", event, order.ID)
	switch event {
	case core.Cooked:
		c.scheduleCleanJob(order)
	case core.Moved:
		c.scheduler.RemoveEntry(order.ID)
		c.scheduleCleanJob(order)
	case core.Delivered:
		c.scheduler.RemoveEntry(order.ID)
	default:
		c.ctx.Log.Errorf("cleaner received an invalid event: %d", event)
		if c.ctx.IsDebug {
			panic("cleaner received an invalid event")
		}
	}
}

func (c *cleaner) scheduleCleanJob(order *core.Order) {
	timeToClean := c.calculateExpireTime(order)
	c.scheduler.AddFuncJob(order.ID, func() {
		// pick food from shelf
		err := c.Kitchen.GetShelf().Pick(order)
		if err != nil {
			if !core.ResourceNotFound.Is(err) {
				c.ctx.Log.WithError(err).Error("unknown error")
			}
			return
		}
		// tell message center this order is discarded
		c.Kitchen.Send(order, core.Discarded)
	}, timeToClean)
}

func (c *cleaner) calculateExpireTime(order *core.Order) time.Time {
	var shelfDecayModifier float64
	if order.ShelfType == core.SingleTempShelf {
		shelfDecayModifier = 1.0
	} else {
		shelfDecayModifier = 2.0
	}
	age := order.RemainLife / (order.DecayRate * shelfDecayModifier)
	return time.Now().Add(time.Second * time.Duration(age))
}

func (c *cleaner) GetOffWork() {
	// block until all scheduled clean job done
	c.scheduler.Stop(false)
	time.Sleep(time.Millisecond * 100)
}
