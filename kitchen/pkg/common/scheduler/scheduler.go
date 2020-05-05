package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	entries *entries
	stop    chan bool
	add     chan *schEntry
	remove  chan EntryID
	running bool

	timer *time.Timer

	mutex     sync.Mutex
	waitGroup sync.WaitGroup
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		entries: newEntries(),
	}
}

func (s *Scheduler) addEntry(entry *schEntry) {
	s.entries.Insert(entry)
}

func (s *Scheduler) removeEntry(id EntryID) {
	s.entries.Remove(id)
}

func (s *Scheduler) updateTimer(now time.Time) {
	top := s.entries.Head()
	if top == nil {
		s.timer.Reset(10000 * time.Hour)
	} else if !top.Next.After(now) {
		s.timer.Reset(now.Add(time.Second).Truncate(time.Second).Sub(now))
	} else {
		s.timer.Reset(top.Next.Sub(now))
	}
}

func (s *Scheduler) runJobs(now time.Time) {
	for {
		top := s.entries.Head()
		if top == nil || top.Next.After(now) {
			break
		}
		s.entries.Pop()
		top.Job.Do()
	}
}

func (s *Scheduler) run() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()

	now := time.Now()
	s.timer = time.NewTimer(10000 * time.Hour)
	for {
		s.updateTimer(now)
		select {
		case now = <-s.timer.C:
			s.runJobs(now)
			if !s.running && s.entries.que.Len() == 0 {
				s.timer.Stop()
				return
			}
		case entry := <-s.add:
			s.addEntry(entry)
			now = time.Now()
		case id := <-s.remove:
			s.removeEntry(id)
			now = time.Now()
		case force := <-s.stop:
			if force || s.entries.que.Len() == 0 {
				s.timer.Stop()
				return
			}
		}
	}
}

func (s *Scheduler) Stop(force bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.running {
		for _, e := range s.entries.que.buf {
			fmt.Println(e)
		}
		s.running = false
		// send stop command
		s.stop <- force
		// block until working thread quit
		s.waitGroup.Wait()
		// close all channels
		close(s.stop)
		close(s.add)
		close(s.remove)
	}
}

func (s *Scheduler) Run() *Scheduler {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.running {
		s.stop = make(chan bool)
		s.add = make(chan *schEntry)
		s.remove = make(chan EntryID)

		go s.run()
		s.running = true
	}
	return s
}

func (s *Scheduler) AddEntry(entry *schEntry) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.running {
		s.add <- entry
	}
}

func (s *Scheduler) AddJob(id EntryID, job Job, next time.Time) {
	s.AddEntry(NewEntry(id, job, next))
}

func (s *Scheduler) RemoveEntry(id EntryID) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.running {
		s.remove <- id
	}
}

type simpleJob struct {
	f func()
}

func (sj *simpleJob) Do() {
	sj.f()
}

func (s *Scheduler) AddFuncJob(id EntryID, f func(), next time.Time) {
	s.AddJob(id, &simpleJob{f: f}, next)
}
