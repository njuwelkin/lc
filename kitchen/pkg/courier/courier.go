package courier

import (
	"math/rand"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/common/worker_pool"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

type courierMgr struct {
	*core.BaseColleague
	ctx      *core.Context
	kitchen  core.Kitchen
	couriers *worker_pool.WorkerPool
}

func NewCourierMgr(ctx *core.Context) *courierMgr {
	return &courierMgr{
		BaseColleague: core.NewBaseColleague(),
		ctx:           ctx,
		couriers:      worker_pool.NewWorkerPool(ctx.NumOfCouriers, 1000).Run(),
	}
}

func (c *courierMgr) Notify(order *core.Order) {
	c.couriers.InsertJob(core.NewJob(func() {
		c.ctx.Log.Info("pick")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(5)))
		order.Status = core.Picked
		c.Kitchen.Send(order)
	}))
}

func (c *courierMgr) GetOffWork() {
	c.couriers.Quit()
}
