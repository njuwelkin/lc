package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	entries *entries
	stop    chan struct{}
	add     chan *schEntry
	remove  chan EntryID
	running bool

	timer *time.Timer
	mutex sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		entries: newEntries(),
		stop:    make(chan struct{}),
		add:     make(chan *schEntry),
		remove:  make(chan EntryID),
	}
}

func (s *Scheduler) addEntry(entry *schEntry) {
	fmt.Printf("%+v\n", entry.Next)
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
	now := time.Now()
	s.timer = time.NewTimer(10000 * time.Hour)
	for {
		s.updateTimer(now)
		select {
		case now = <-s.timer.C:
			s.runJobs(now)
		case entry := <-s.add:
			s.addEntry(entry)
			now = time.Now()
		case id := <-s.remove:
			s.removeEntry(id)
			now = time.Now()
		case <-s.stop:
			s.timer.Stop()
			break
		}
	}
}

func (s *Scheduler) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.running {
		s.stop <- struct{}{}
		s.running = false
	}
}

func (s *Scheduler) Run() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.running {
		go s.run()
		s.running = true
	}
}

func (s *Scheduler) AddEntry(entry *schEntry) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.running {
		s.add <- entry
	} else {
		s.addEntry(entry)
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
	} else {
		s.removeEntry(id)
	}
}
