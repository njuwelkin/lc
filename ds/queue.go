package ds

import (
	"errors"
)

type ContentType = interface{}

type Queue struct {
	buf        []ContentType
	head, tail int
}

func NewQueue() *Queue {
	return &Queue{make([]ContentType, 32), 0, 0}
}

func (q *Queue) Add(val ContentType) {
	if (q.tail+1)%len(q.buf) == q.head {
		origSize := len(q.buf)
		q.buf = append(q.buf, make([]ContentType, origSize)...)
		if q.tail < q.head {
			for i := 0; i < q.head; i++ {
				q.buf[origSize+i] = q.buf[i]
			}
			q.tail += origSize
		}
	}
	q.buf[q.tail] = val
	q.tail = (q.tail + 1) % len(q.buf)
}

func (q *Queue) Delete() (ContentType, error) {
	if q.tail == q.head {
		return 0, errors.New("empty")
	}
	ret := q.buf[q.head]
	q.head = (q.head + 1) % len(q.buf)
	return ret, nil
}

func (q *Queue) Head() (ContentType, error) {
	if q.tail == q.head {
		return 0, errors.New("empty")
	}
	return q.buf[q.head], nil
}

func (q *Queue) IsEmpty() bool {
	return q.tail == q.head
}
