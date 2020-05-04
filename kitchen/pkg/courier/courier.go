package courier

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/common/worker_pool"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
)

const (
	minPickDuration = 2
	maxPickDuration = 6
)

type courierMgr struct {
	*core.BaseColleague
	ctx      *core.Context
	couriers *worker_pool.WorkerPool

	order2courier map[string]*courierJob
	o2cMutex      sync.Mutex

	ongoingJob int64
	pendingJob int64
}

func NewCourierMgr(ctx *core.Context) *courierMgr {
	return &courierMgr{
		BaseColleague: core.NewBaseColleague(),
		ctx:           ctx,
		couriers:      worker_pool.NewWorkerPool(ctx.NumOfCouriers, 4096).Run(),
		order2courier: map[string]*courierJob{},
		ongoingJob:    0,
		pendingJob:    0,
	}
}

func (c *courierMgr) Notify(order *core.Order, event core.Event) {
	c.ctx.Log.Infof("courierMgr: receive Event %d for order %s", event, order.ID)
	switch event {
	case core.Accept:
		// accept a new order, dispatch a courier to pickup
		job := c.addJob(order)
		order.EstimatePickTime = c.latestPickTime()
		atomic.AddInt64(&c.pendingJob, 1)
		c.couriers.InsertJob(job)
	case core.Discarded:
		// a order has been discarded, tell courier to abort picking
		c.o2cMutex.Lock()
		job, found := c.order2courier[order.ID]
		if !found {
			c.o2cMutex.Unlock()
			return
		}
		job.cancel <- struct{}{}
		c.o2cMutex.Unlock()
	default:
		// this will not happen
		c.ctx.Log.Errorf("invalid envent %d", event)
		if c.ctx.IsDebug {
			panic("invalid enven")
		}
	}
}

func (c *courierMgr) GetOffWork() {
	c.couriers.Quit()
}

func (c *courierMgr) latestPickTime() time.Time {
	inSecond := int64(maxPickDuration) + // complete ongoing job
		(c.pendingJob/int64(c.ctx.NumOfCouriers))*maxPickDuration + // complete pending job
		int64(maxPickDuration) // time to arrive at kitchen
	return time.Now().Add(time.Second * time.Duration(inSecond))
}

func (c *courierMgr) addJob(order *core.Order) *courierJob {
	c.o2cMutex.Lock()
	defer c.o2cMutex.Unlock()

	job := newCourierJob(c, order)
	c.order2courier[order.ID] = job
	return job
}

func (c *courierMgr) removeJob(id string) {
	c.o2cMutex.Lock()
	defer c.o2cMutex.Unlock()

	job := c.order2courier[id]
	close(job.cancel)
	delete(c.order2courier, id)
}

//
type courierJob struct {
	mgr   *courierMgr
	order *core.Order

	cancel chan struct{}
}

func newCourierJob(mgr *courierMgr, order *core.Order) *courierJob {
	return &courierJob{
		mgr:    mgr,
		order:  order,
		cancel: make(chan struct{}, 1),
	}
}

func (cj *courierJob) Do() {
	cj.mgr.ctx.Log.Infof("courier take order %+v", cj.order)
	atomic.AddInt64(&cj.mgr.pendingJob, -1)
	atomic.AddInt64(&cj.mgr.ongoingJob, 1)
	defer atomic.AddInt64(&cj.mgr.ongoingJob, -1)
	defer cj.mgr.removeJob(cj.order.ID)

	// go to kitchen
	if !cj.gotoKitchen(cj.order) {
		return
	}
	// pick the food
	err := cj.mgr.Kitchen.GetShelf().Pick(cj.order)
	if err != nil {
		cj.mgr.ctx.Log.WithError(err).Warnf("order %s not found", cj.order.ID)
		return
	}
	// update order status and notify msg center
	cj.order.Status = "delivered"
	cj.mgr.ctx.Log.Infof("Delivered order %+v", cj.order)
	cj.mgr.Kitchen.Send(cj.order, core.Delivered)
}

func (cj *courierJob) gotoKitchen(order *core.Order) bool {
	timeOnTheWay := time.Second * time.Duration(minPickDuration+rand.Intn(maxPickDuration-minPickDuration+1))
	// give a accurate estimation of pick time
	order.EstimatePickTime = time.Now().Add(timeOnTheWay)

	// impossible to arrive at kitchen in time, abort this order
	if value := order.EstimatePickValue(false); value <= 0 {
		cj.mgr.ctx.Log.Infof("estimateValue: %d", value)
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
	case <-cj.cancel:
		// order is discarded, terminate the task
		cj.mgr.ctx.Log.Infof("order %s is discarded, abort picking job", cj.order.ID)
		return false
	}
}
