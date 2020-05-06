package worker_pool_test

import (
	. "github.com/onsi/gomega"
	"sync/atomic"
	"testing"
	"time"

	. "kitchen/pkg/common/worker_pool"
	"kitchen/pkg/test"
)

var count int64 = 0

type job struct {
}

func (j *job) Do() {
	count++
}

func TestWorkerPool(tt *testing.T) {
	t := test.NewTest(tt)
	wp := NewWorkerPool(5, 5)

	t.Run("add job before running worker pool", func() {
		wp.InsertFuncJob(func() {
			count++
		})
		time.Sleep(time.Second)
		// worker pool is not running, nothing will happen
		t.Expect(int(count)).To(Equal(0))
	})

	t.Run("add job when worker pool is running", func() {
		count = 0
		wp.Run()
		for i := 0; i < 5; i++ {
			wp.InsertFuncJob(func() {
				time.Sleep(time.Second)
				atomic.AddInt64(&count, 1)
			})
		}
		t.Expect(int(count)).To(Equal(0))
		// wait until all jobs done
		wp.Quit()
		t.Expect(int(count)).To(Equal(5))
	})

	t.Run("amount of jobs more than workers", func() {
		count = 0
		wp.Run()
		for i := 0; i < 8; i++ {
			wp.InsertFuncJob(func() {
				time.Sleep(time.Second)
				atomic.AddInt64(&count, 1)
			})
		}
		t.Expect(int(count)).To(Equal(0))
		time.Sleep(time.Millisecond * 1100)
		t.Expect(int(count)).To(Equal(5))
		// wait until all jobs done
		wp.Quit()
		t.Expect(int(count)).To(Equal(8))
	})

	t.Run("amount of jobs more than workers plus size of buffer", func() {
		count = 0
		wp.Run()
		for i := 0; i < 15; i++ {
			// it will block on 11th job
			// the loop won't be ended until the first 5 jobs done
			wp.InsertFuncJob(func() {
				time.Sleep(time.Second)
				atomic.AddInt64(&count, 1)
			})
		}

		// first 5 jobs already done
		t.Expect(int(count)).To(Equal(5))
		time.Sleep(time.Millisecond * 1100)
		t.Expect(int(count)).To(Equal(10))
		// wait until all jobs done
		wp.Quit()
		t.Expect(int(count)).To(Equal(15))
	})

}
