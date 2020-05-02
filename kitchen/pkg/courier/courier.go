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
	order.EstimatePickTime = time.Now().Add(time.Hour * 100)
	c.couriers.InsertFuncJob(func() {
		// go to kitchen
		if !gotoKitchen(order) {
			return
		}
		// pick the food
		c.Kitchen.GetShelf().Pick(order.ID)
		// update order status and notify msg center
		//c.ctx.Log.Info("picked")
		order.Status = "delivered"
		c.Kitchen.Send(order, core.Delivered)
	})
}

func (c *courierMgr) GetOffWork() {
	c.couriers.Quit()
}

func gotoKitchen(order *core.Order) bool {
	timeOnTheWay := time.Second * time.Duration(2+rand.Intn(5))
	order.EstimatePickTime = time.Now().Add(timeOnTheWay)
	timer := time.NewTimer(timeOnTheWay)
	defer timer.Stop()

	select {
	case <-timer.C:
		// arrive at the kitchen
		return true
	case <-order.IsCanceled:
		// order is discarded, terminate the task
		return false
	}
}
