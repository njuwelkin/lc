package courier

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/common/worker_pool"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

type courierMgr struct {
	*core.BaseColleague
	ctx      *core.Context
	couriers *worker_pool.WorkerPool

	ongoingJob int64
	pendingJob int64
}

func NewCourierMgr(ctx *core.Context) *courierMgr {
	return &courierMgr{
		BaseColleague: core.NewBaseColleague(),
		ctx:           ctx,
		couriers:      worker_pool.NewWorkerPool(ctx.NumOfCouriers, 4096).Run(),
		ongoingJob:    0,
		pendingJob:    0,
	}
}

func (c *courierMgr) Notify(order *core.Order, event core.Event) {
	c.ctx.Log.Infof("courierMgr: receive Event %d for order %s", event, order.ID)
	switch event {
	case core.Accepted:
		// accept a new order, dispatch a courier to pickup
		job := newCourierJob(c, order)
		order.EstimatePickTime = c.latestPickTime()
		atomic.AddInt64(&c.pendingJob, 1)
		c.couriers.InsertJob(job)
	default:
		c.ctx.Log.Errorf("courierMgr received an invalid envent %d", event)
		if c.ctx.IsDebug {
			panic("invalid enven")
		}
	}
}

func (c *courierMgr) GetOffWork() {
	c.couriers.Quit()
	time.Sleep(time.Millisecond * 100)
}

func (c *courierMgr) latestPickTime() time.Time {
	inSecond := int64(c.ctx.MaxPickDuration) + // complete ongoing job
		(c.pendingJob/int64(c.ctx.NumOfCouriers))*int64(c.ctx.MaxPickDuration) + // complete pending job
		int64(c.ctx.MaxPickDuration) // time to arrive at kitchen
	return time.Now().Add(time.Second * time.Duration(inSecond))
}

type courierJob struct {
	mgr   *courierMgr
	order *core.Order
}

func newCourierJob(mgr *courierMgr, order *core.Order) *courierJob {
	return &courierJob{
		mgr:   mgr,
		order: order,
	}
}

func (cj *courierJob) Do() {
	cj.mgr.ctx.Log.Infof("courier take order %+v", cj.order)
	atomic.AddInt64(&cj.mgr.pendingJob, -1)
	atomic.AddInt64(&cj.mgr.ongoingJob, 1)
	defer atomic.AddInt64(&cj.mgr.ongoingJob, -1)

	// go to kitchen
	if !cj.gotoKitchen(cj.order) {
		return
	}

	// wait until food is ready
	<-cj.order.Ready

	// pick the food
	err := cj.mgr.Kitchen.GetShelf().Pick(cj.order)
	if err != nil {
		if !core.ResourceNotFound.Is(err) {
			cj.mgr.ctx.Log.WithError(err).Warnf("unknown error")
		}
		cj.mgr.ctx.Log.Infof("order %s not exist", cj.order.ID)
		return
	}

	// update order status and notify msg center
	cj.order.Status = "picked"
	cj.mgr.Kitchen.Send(cj.order, core.Picked)
	// cj.deliever(order)
	cj.order.Status = "delivered"
	cj.mgr.ctx.Log.Infof("Delivered order %+v", cj.order)
	cj.mgr.Kitchen.Send(cj.order, core.Delivered)
}

func (cj *courierJob) gotoKitchen(order *core.Order) bool {
	minPickDuration := cj.mgr.ctx.MinPickDuration
	maxPickDuration := cj.mgr.ctx.MaxPickDuration
	timeOnTheWay := time.Second * time.Duration(minPickDuration+rand.Intn(maxPickDuration-minPickDuration+1))
	// give a accurate estimation of pick time
	order.EstimatePickTime = time.Now().Add(timeOnTheWay)

	// impossible to arrive at kitchen in time, abort this order
	if value := order.EstimatePickValue(false); value <= 0 {
		cj.mgr.ctx.Log.Infof("estimateValue: %f", value)
		cj.mgr.ctx.Log.Warnf("impossible to pick order %s in time, abort", cj.order.ID)
		cj.order.EstimatePickTime = time.Now().Add(time.Hour * 10000)
		return false
	}

	timer := time.NewTimer(timeOnTheWay)
	defer timer.Stop()

	// go to kitchen, and stop if receive the cancel command
	select {
	case <-timer.C:
		// arrive at the kitchen
		return true
	case <-order.Cancel:
		// order is discarded, terminate the task
		cj.mgr.ctx.Log.Infof("order %s is discarded, abort picking job", cj.order.ID)
		return false
	}
}
