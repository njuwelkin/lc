package kitchen

import (
	. "github.com/onsi/gomega"
	"sync/atomic"
	"testing"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/core"
	. "github.com/njuwelkin/lc/kitchen/pkg/kitchen"
	"github.com/njuwelkin/lc/kitchen/pkg/test"
)

var orders []*core.Order = [][]*core.Order {
	&core.Order{
		ID: "1",
		Temp: core.Hot,

	}
}

func TestShelfSet(tt *testing.T) {
	t := test.NewTest(tt)
	t.ShelfCap.Hot = 1
	t.ShelfCap.Cold = 1
	t.ShelfCap.Frozen = 1
	t.ShelfCap.Overflow = 2

	shelf := newShelfSet(t.Context)

	t.Run("put food on single temp shelf", func() {
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
