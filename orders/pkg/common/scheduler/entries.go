package scheduler

import (
	"container/heap"
	//"fmt"
	"time"
)

type EntryID string

type Job interface {
	Do()
}

type schEntry struct {
	ID  EntryID
	Job Job

	Next time.Time
}

func NewEntry(id EntryID, job Job, next time.Time) *schEntry {
	return &schEntry{
		ID:   id,
		Job:  job,
		Next: next,
	}
}

type entryQue struct {
	buf   []*schEntry
	idMap map[EntryID]int
}

func (q *entryQue) Len() int {
	return len(q.buf)
}

func (q *entryQue) Less(i, j int) bool {
	return !q.buf[i].Next.After(q.buf[j].Next)
}

func (q *entryQue) Swap(i, j int) {
	q.idMap[q.buf[i].ID] = j
	q.idMap[q.buf[j].ID] = i
	q.buf[i], q.buf[j] = q.buf[j], q.buf[i]
}

func (q *entryQue) Push(x interface{}) {
	entry := x.(*schEntry)
	q.buf = append(q.buf, entry)
	q.idMap[entry.ID] = len(q.buf) - 1
}

func (e *entryQue) Pop() interface{} {
	ret := e.buf[e.Len()-1]
	delete(e.idMap, ret.ID)
	e.buf = e.buf[:e.Len()-1]
	return ret
}

type entries struct {
	que entryQue
}

func newEntries() *entries {
	ret := entries{}
	ret.que = entryQue{
		buf:   []*schEntry{},
		idMap: map[EntryID]int{},
	}
	return &ret
}

func (e *entries) Insert(entry *schEntry) {
	heap.Push(&e.que, entry)
}

func (e *entries) Head() *schEntry {
	if e.que.Len() == 0 {
		return nil
	}
	return e.que.buf[0]
}

func (e *entries) Pop() *schEntry {
	x := heap.Pop(&e.que)
	entry := x.(*schEntry)
	return entry
}

func (e *entries) Remove(id EntryID) {
	i, found := e.que.idMap[id]
	if found {
		delete(e.que.idMap, id)
		heap.Remove(&e.que, i)
	}
}

func (e *entries) Get(id EntryID) *schEntry {
	i := e.que.idMap[id]
	if i >= e.que.Len() {
		return nil
	}
	return e.que.buf[i]
}
