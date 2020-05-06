package scheduler_test

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"

	. "kitchen/pkg/common/scheduler"
	"kitchen/pkg/test"
)

var count int = 0

type job struct {
}

func (j *job) Do() {
	count++
}

func TestScheduler(tt *testing.T) {
	t := test.NewTest(tt)
	sch := NewScheduler()

	t.Run("add job before running scchedler", func() {
		now := time.Now()
		sch.AddJob("1", &job{}, now.Add(time.Second))
		time.Sleep(time.Second)
		// scheduler is not running, nothing will happen
		t.Expect(count).To(Equal(0))
	})

	t.Run("add job when scchedler is running", func() {
		count = 0
		sch.Run()
		sch.AddJob("1", &job{}, time.Now().Add(time.Second))
		t.Expect(count).To(Equal(0))
		time.Sleep(time.Millisecond * 1100)
		t.Expect(count).To(Equal(1))
	})

	t.Run("remove a job", func() {
		count = 0
		sch.AddJob("1", &job{}, time.Now().Add(10*time.Second))
		t.Expect(count).To(Equal(0))
		sch.RemoveEntry("1")
		time.Sleep(time.Millisecond * 1100)
		t.Expect(count).To(Equal(0))
	})

	t.Run("stop the schedluer", func() {
		count = 0
		now := time.Now()
		sch.AddJob("1", &job{}, now.Add(time.Millisecond*100))
		sch.AddJob("2", &job{}, now.Add(time.Millisecond*500))
		sch.AddJob("3", &job{}, now.Add(time.Millisecond*900))
		t.Expect(count).To(Equal(0))
		sch.Stop(false)
		// make sure all jobs done
		t.Expect(count).To(Equal(3))
	})
	t.Run("force stop a schedluer", func() {
		count = 0
		now := time.Now()
		sch.Run()
		sch.AddJob("1", &job{}, now.Add(time.Millisecond*100))
		sch.AddJob("2", &job{}, now.Add(time.Millisecond*500))
		sch.AddJob("3", &job{}, now.Add(time.Millisecond*900))
		t.Expect(count).To(Equal(0))
		sch.Stop(true)
		// force stopped, jobs are not done
		t.Expect(count).To(Equal(0))
	})

}
