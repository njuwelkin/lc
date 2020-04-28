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
	c.couriers.InsertFuncJob(func() {
		// go to kitchen
		c.ctx.Log.Info("picking")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(5)))
		// pick the food
		//c.Kitchen.GetShelf().Pick(order.ID)
		// update order status and notify msg center
		c.ctx.Log.Info("picked")
		order.Status = core.Delivered
		c.Kitchen.Send(order)
		count++
	})
}

func (c *courierMgr) GetOffWork() {
	c.couriers.Quit()
}
