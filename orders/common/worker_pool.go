package utils

import ()

type JobStatus int

const (
	JobStatusWaiting JobStatus = iota
	JobStatusRunning
	JobStatusDone
)

type Job interface {
	Do()
}

type statusJob interface {
	Job
	SetStatus(status JobStatus)
}

type WorkerPool struct {
	jobChannel  chan Job
	quitChannel chan bool
	workers     int
	chanSize    int
}

func NewWorkerPool(workers int, chanSize int) *WorkerPool {
	return &WorkerPool{
		jobChannel:  make(chan Job, chanSize),
		quitChannel: make(chan bool),
		workers:     workers,
		chanSize:    chanSize,
	}
}

func (workerPool *WorkerPool) execute() {
	for {
		select {
		case job := <-workerPool.jobChannel:
			if sj, ok := job.(statusJob); ok {
				sj.SetStatus(JobStatusRunning)
			}

			job.Do()

			if sj, ok := job.(statusJob); ok {
				sj.SetStatus(JobStatusDone)
			}
		case <-workerPool.quitChannel:
			break
		}
	}
}

func (workerPool *WorkerPool) Run() *WorkerPool {
	for i := 0; i < workerPool.workers; i++ {
		go workerPool.execute()
	}
	return workerPool
}

func (workerPool *WorkerPool) InsertJob(job Job) {
	if sj, ok := job.(statusJob); ok {
		sj.SetStatus(JobStatusWaiting)
	}

	workerPool.jobChannel <- job
}

func (workerPool *WorkerPool) Quit() {
	for i := 0; i < workerPool.workers; i++ {
		workerPool.quitChannel <- true
	}
}
